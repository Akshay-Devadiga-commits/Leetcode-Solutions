type ProductOfNumbers struct {
    // Stores cumulative product of the stream
    prefixProduct []int
    size          int
}

func Constructor() ProductOfNumbers {
    // Initialize the product list with 1 to handle multiplication logic
    return ProductOfNumbers{
        prefixProduct: []int{1},
        size:          0,
    }
}

func (p *ProductOfNumbers) Add(num int) {
    if num == 0 {
        // If num is 0, reset the cumulative products since multiplication
        // with 0 invalidates previous products
        p.prefixProduct = []int{1}
        p.size = 0
    } else {
        // Append the cumulative product of the current number with the last
        // product
        p.prefixProduct = append(p.prefixProduct, p.prefixProduct[p.size]*num)
        p.size++
    }
}

func (p *ProductOfNumbers) GetProduct(k int) int {
    // Check if the requested product length exceeds the size of the valid
    // product list
    if k > p.size {
        return 0
    }

    // Compute the product of the last k elements using division
    return p.prefixProduct[p.size] / p.prefixProduct[p.size-k]
}