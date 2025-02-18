// Package domain_v2 is updated version of domain package
package domain_v2

import "gokit/examples/generic-alias/domain"

type Customer = domain.Customer

type User[T any] = domain.User[T]
