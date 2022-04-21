package handlers

import (
    "errors"
    "github.com/arcs/pastee/utl"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/gin/binding"
    "net/http"
)

const MaxPasteCharacters = 2 * 1000 * 1000

type PastePostRequest struct {
    Content string `json:"content"`
    Encrypt bool   `json:"encrypt"`
}

type PastePostResponse struct {
    Id  string `json:"id"`
    Key string `json:"key"`
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
    key := c.Query("key")

    content := paste.Content
    if paste.Hashed {
        decrypted, err := utl.AesDecryptCBC(paste.Content, key)
        if err == nil {
            content = decrypted
        }
    }

    c.String(http.StatusOK, "%s", content)
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

    key := ""
    if pastePostRequest.Encrypt {
        key = utl.GenerateToken(32)
        pastePostRequest.Content, err = utl.AesEncryptCBC(pastePostRequest.Content, key)
        if err != nil {
            c.AbortWithError(http.StatusInternalServerError, errors.New("error saving entry"))
            return
        }
    }
    id := utl.GenerateToken(8)
    encrypted := key != ""
    err = s.store.SavePaste(id, pastePostRequest.Content, encrypted)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, errors.New("error saving entry"))
        return
    }
    response := PastePostResponse{
        Id:  id,
        Key: key,
    }
    c.JSON(http.StatusOK, response)
}
