package handlers

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    "go.mongodb.org/mongo-driver/v2/bson"

    "std-mongo/database"
    "std-mongo/models"
)

func ProductRoutes(mux *http.ServeMux) {
    mux.HandleFunc("GET /products", getProducts)
    mux.HandleFunc("GET /products/{id}", getProduct)
    mux.HandleFunc("POST /products", createProduct)
    mux.HandleFunc("PUT /products/{id}", updateProduct)
    mux.HandleFunc("DELETE /products/{id}", deleteProduct)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product
    coll := database.GetCollection("products")
    cursor, _ := coll.Find(context.TODO(), bson.M{})
    cursor.All(context.TODO(), &products)
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    objID, _ := bson.ObjectIDFromHex(id)
    
    var product models.Product
    coll := database.GetCollection("products")
    if err := coll.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&product); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Product not found",
        })
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Invalid input",
        })
        return
    }
    product.CreatedAt = time.Now()
    product.UpdatedAt = time.Now()

    coll := database.GetCollection("products")
    result, _ := coll.InsertOne(context.TODO(), product)
    product.ID = result.InsertedID.(bson.ObjectID)
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(product)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    objID, _ := bson.ObjectIDFromHex(id)
    
    var updatedProduct models.Product
    if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Invalid input",
        })
        return
    }
    updatedProduct.UpdatedAt = time.Now()
    updatedProduct.ID = objID
    update := bson.M{
        "$set": bson.M{
            "name":        updatedProduct.Name,
            "description": updatedProduct.Description,
            "price":       updatedProduct.Price,
            "updated_at":  updatedProduct.UpdatedAt,
        },
    }
    
    coll := database.GetCollection("products")
    if result, _ := coll.UpdateOne(context.TODO(), bson.M{"_id": objID}, update); result.MatchedCount == 0 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Product not found",
        })
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedProduct)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    objID, _ := bson.ObjectIDFromHex(id)

    coll := database.GetCollection("products")
    if result, _ := coll.DeleteOne(context.TODO(), bson.M{"_id": objID}); result.DeletedCount == 0 {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(map[string]string{
            "error": "Product not found",
        })
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Product deleted successfully",
    })
}
