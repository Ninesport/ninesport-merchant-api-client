package client

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/google/go-querystring/query"
)

func ToQueryString(obj any, keyDesc bool) (string, error) {
	values, err := query.Values(obj)
	if err != nil {
		return "", err
	}

	var keys []string
	for k := range values {
		keys = append(keys, k)
	}

	if keyDesc {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	var strs []string
	for _, key := range keys {
		value := values.Get(key)
		strs = append(strs, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(strs, "&"), nil
}

func Sign(secretKey string, queryString string) (string, error) {
	signString := fmt.Sprintf("%s&secretKey=%s", queryString, secretKey)
	if DEBUG {
		fmt.Printf("[signString]: %s\n", signString)
	}
	h := sha256.New()
	_, err := h.Write([]byte(signString))
	if err != nil {
		return "", err
	}
	sign := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	if DEBUG {
		fmt.Printf("[sign]: %s\n", sign)
	}
	return sign, nil
}
