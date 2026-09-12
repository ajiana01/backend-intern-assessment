package main

import (
	"errors"
	"fmt"
)

type CartItem struct {
	ProductID string
	Name      string
	Price     float64
	Quantity  int
}

type Voucher struct {
	Code            string
	DiscountPercent float64
	MaxDiscount     float64
	MinPurchase     float64
}

func CalculateFinalPrice(items []CartItem, voucher *Voucher) (subtotal float64, discount float64, total float64, err error) {

	if len(items) == 0 {
		return 0, 0, 0, errors.New("cart cannot be empty")
	}

	for _, item := range items {

		if item.Quantity <= 0 || item.Price < 0 {
			return 0, 0, 0, errors.New("invalid item price or quantity")
		}

		subtotal += item.Price * float64(item.Quantity)
	}

	if voucher != nil && subtotal >= voucher.MinPurchase {

		discount = subtotal * (voucher.DiscountPercent / 100)

		if discount > voucher.MaxDiscount {
			discount = voucher.MaxDiscount
		}
	}

	total = subtotal - discount
	return subtotal, discount, total, nil
}

func main() {

	// Case 1: Succesful discount
	cartItems1 := []CartItem{
		{ProductID: "MS", Name: "Mouse", Price: 100_000, Quantity: 2},
	}

	voucher1 := &Voucher{
		Code:            "DISC10",
		DiscountPercent: 10,
		MaxDiscount:     50_000,
		MinPurchase:     100_000,
	}

	subtotal1, discount1, total1, err1 := CalculateFinalPrice(cartItems1, voucher1)

	if err1 != nil {
		fmt.Println("Error:", err1.Error())
	} else {
		fmt.Println("\nCase 1: Successful discount")
		fmt.Println("Subtotal:", subtotal1)
		fmt.Println("Discount:", discount1)
		fmt.Println("Total:", total1)
	}

	// Case 2: No discount applied
	cartItems2 := []CartItem{
		{ProductID: "KB", Name: "Keyboard", Price: 80_000, Quantity: 1},
	}

	subtotal2, discount2, total2, err2 := CalculateFinalPrice(cartItems2, nil)

	if err2 != nil {
		fmt.Println("Error:", err2.Error())
	} else {
		fmt.Println("\nCase 2: No discount applied")
		fmt.Println("Subtotal:", subtotal2)
		fmt.Println("Discount:", discount2)
		fmt.Println("Total:", total2)
	}

	// Case 3: Empty cart
	cartItems3 := []CartItem{}
	subtotal3, discount3, total3, err3 := CalculateFinalPrice(cartItems3, nil)

	if err3 != nil {
		fmt.Println("\nCase 3: Empty cart")
		fmt.Println("Error:", err3.Error())
	} else {
		fmt.Println("\nCase 3: Empty cart")
		fmt.Println("Subtotal:", subtotal3)
		fmt.Println("Discount:", discount3)
		fmt.Println("Total:", total3)
	}

	// Case 4: Minimum purchase not met
	cartItems4 := []CartItem{
		{ProductID: "HD", Name: "Headphones", Price: 50_000, Quantity: 1},
	}
	subtotal4, discount4, total4, err4 := CalculateFinalPrice(cartItems4, voucher1)

	if err4 != nil {
		fmt.Println("Error:", err4.Error())
	} else {
		fmt.Println("\nCase 4: Minimum purchase not met")
		fmt.Println("Subtotal:", subtotal4)
		fmt.Println("Discount:", discount4)
		fmt.Println("Total:", total4)
	}

	// Case 5: Invalid item price or quantity
	cartItems5 := []CartItem{
		{ProductID: "MS", Name: "Mouse", Price: -100_000, Quantity: 2},
	}
	subtotal5, discount5, total5, err5 := CalculateFinalPrice(cartItems5, voucher1)

	if err5 != nil {
		fmt.Println("\nCase 5: Invalid item price or quantity")
		fmt.Println("Error:", err5.Error())
	} else {
		fmt.Println("\nCase 5: Invalid item price or quantity")
		fmt.Println("Subtotal:", subtotal5)
		fmt.Println("Discount:", discount5)
		fmt.Println("Total:", total5)
	}
}
