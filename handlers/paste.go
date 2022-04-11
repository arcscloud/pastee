package handlers

import (
    "errors"
    "github.com/arcs/pastee/rand"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "log"
    "net/http"
)

const MaxPasteCharacters = 2 * 1000 * 1000

type PastePostRequest struct {
    Content string `json:"content"`
    Hash    bool   `json:"hash"`
}

type PastePostResponse struct {
    Id   string `json:"id"`
    Hash string `json:"hash"`
}

func (s defaultServer) pasteOptions() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, PUT")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
    }
}

func (s defaultServer) pasteGet(c *gin.Context) {
    id := c.Param("id")
    paste, err := s.store.GetPaste(id)
    if err != nil {
        c.String(http.StatusNotFound, "%s", "paste not found")
        return
    }
    hash := c.Query("hash")
    if paste.Hash != "" && hash != paste.Hash {
        c.String(http.StatusNotFound, "%s", "paste not found")
        return
    }

    c.String(http.StatusOK, "%s", paste.Content)
}

func (s defaultServer) pastePost(c *gin.Context) {
    pastePostRequest := new(PastePostRequest)
    err := c.MustBindWith(pastePostRequest, binding.JSON)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"response": "Malformed JSON"})
        return
    }
    if pastePostRequest.Content == "" {
        c.JSON(http.StatusBadRequest, gin.H{"response": "Paste can not be empty!"})
        return
    }

    hash := ""
    if pastePostRequest.Hash {
        hash = rand.GenerateToken(32)
    }
    id := rand.GenerateToken(8)
    err = s.store.SavePaste(id, hash, pastePostRequest.Content)
    if err != nil {
        log.Printf("Error saving entry: %v\n", err)
        c.AbortWithError(http.StatusInternalServerError, errors.New("error saving entry"))
        return
    }
    response := PastePostResponse{
        Id:   id,
        Hash: hash,
    }
    c.JSON(http.StatusOK, response)
}
