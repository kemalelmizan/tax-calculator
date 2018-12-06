package controller

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestPingHealthy(t *testing.T) {
	Healthy = 1
	handler := Ping()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusOK, status)
	}
}

func TestPingUnhealthy(t *testing.T) {
	Healthy = 0
	handler := Ping()

	req, err := http.NewRequest("GET", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusServiceUnavailable, status)
	}
}

func TestPingWrongMethod(t *testing.T) {
	Healthy = 1
	handler := Ping()

	req, err := http.NewRequest("POST", "/ping", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusNotFound, status)
	}
}

func TestIndex(t *testing.T) {
	handler := Index()

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusNotFound, status)
	}
}

func TestGetBill(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM products (.+)").
		WithArgs("").
		WillReturnRows(sqlmock.NewRows([]string{"name", "tax_code", "price"}).
			AddRow("Lucky Stretch", 2, 110000))

	got := GetBill(db, log.New(os.Stdout, "test: ", log.LstdFlags))

	req, err := http.NewRequest("GET", "/bill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := got

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusOK, status)
	}

	expected := "{\"success\":true,\"data\":{\"products\":[{\"name\":\"Lucky Stretch\",\"tax_code\":2,\"type\":\"Tobacco\",\"refundable\":\"no\",\"price\":\"1100.00\",\"tax\":\"32.00\",\"amount\":\"1132.00\"}],\"total\":\"1132.00\"},\"error_message\":\"\"}\n"
	require.Equal(t, expected, rr.Body.String())
}

func TestBillWrongMethod(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	handler := GetBill(db, log.New(os.Stdout, "test: ", log.LstdFlags))

	req, err := http.NewRequest("POST", "/bill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusNotFound, status)
	}
}

func TestGetBillError(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT (.+) FROM products (.+)").
		WithArgs("").WillReturnError(errors.New("a"))

	got := GetBill(db, log.New(os.Stdout, "test: ", log.LstdFlags))

	req, err := http.NewRequest("GET", "/bill", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := got

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusInternalServerError, status)
	}

	expected := "{\"success\":false,\"data\":null,\"error_message\":\"a\"}\n"
	require.Equal(t, expected, rr.Body.String())
}

func TestPostProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("INSERT INTO products (.+)").
		WithArgs("Lucky Strike", 2, 110000).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(1))

	got := PostProducts(db, log.New(os.Stdout, "test: ", log.LstdFlags))

	req, err := http.NewRequest("POST", "/products", bytes.NewBuffer([]byte(`{"data":[{"name":"Lucky Strike","tax_code":2,"price":110000}]}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := got
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusOK, status)
	}

	expected := "{\"success\":true,\"data\":null,\"error_message\":\"\"}\n"
	require.Equal(t, expected, rr.Body.String())
}

func TestPostProductsWrongMethod(t *testing.T) {
	db, _, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	handler := PostProducts(db, log.New(os.Stdout, "test: ", log.LstdFlags))

	req, err := http.NewRequest("PATCH", "/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		require.Equal(t, http.StatusNotFound, status)
	}
}
