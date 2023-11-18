package handler

import (
	"github.com/akmyrzza/yandex-url/internal/shortener"
	"github.com/akmyrzza/yandex-url/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_ShortURL(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		testURL     string
	}

	tests := []struct {
		name          string
		h             Handler
		request       string
		requestMethod string
		want          want
	}{
		{
			name:          "1st test",
			request:       "/",
			requestMethod: http.MethodPost,
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
				testURL:     "www.example.com",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := gin.Default()
			r.POST(tt.request, tt.h.ShortURL)

			newStorage := *storage.NewStorage()
			newService := shortener.NewShortener(newStorage)
			tt.h.service = *newService
			request := httptest.NewRequest(tt.requestMethod, tt.request, strings.NewReader(tt.want.testURL))

			w := httptest.NewRecorder()
			r.ServeHTTP(w, request)
			//tt.rest.HShortenerURL(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			reqObj, err := io.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			shortURL := string(reqObj)

			parts := strings.Split(shortURL, "/")

			originalURL, exist := newStorage.StorageURL[parts[len(parts)-1]]
			if !exist {
				require.Fail(t, "Expected short URL in urlMap", newStorage.StorageURL)
			}

			assert.Equal(t, tt.want.testURL, originalURL)
		})
	} //
}

func TestHandler_OriginalURL(t *testing.T) {
	type want struct {
		location   string
		statusCode int
	}

	tests := []struct {
		name          string
		h             Handler
		requestMethod string
		want          want
	}{
		{
			name:          "1st test",
			h:             Handler{},
			requestMethod: http.MethodGet,
			want: want{
				location:   "https://www.youtube.com",
				statusCode: http.StatusTemporaryRedirect,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.GET("/:id", tt.h.OriginalURL)

			testShortURL := "Abcdefgh"

			newStorage := *storage.NewStorage()
			newStorage.StorageURL[testShortURL] = tt.want.location

			newService := shortener.NewShortener(newStorage)
			tt.h.service = *newService

			request := httptest.NewRequest(tt.requestMethod, "/"+testShortURL, nil)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, request)

			result := w.Result()
			err := result.Body.Close()
			require.NoError(t, err)

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.location, result.Header.Get("Location"))
		})
	}
}
