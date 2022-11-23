package models

type ThreeAddressCode interface {
	GetThreeAddressCode() ThreeAddressCode
	ToStringCode() string
}
