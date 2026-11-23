type ProductOfNumbers struct {
    prefixProduct []int
    size          int
}

func Constructor() ProductOfNumbers {
    return ProductOfNumbers{
        prefixProduct: []int{1},
        size:          0,
    }
}

func (p *ProductOfNumbers) Add(num int) {
    if num == 0 {
        p.prefixProduct = []int{1}
        p.size = 0
    } else {
        p.prefixProduct = append(p.prefixProduct, p.prefixProduct[p.size]*num)
        p.size++
    }
}

func (p *ProductOfNumbers) GetProduct(k int) int {
    if k > p.size {
        return 0
    }
    return p.prefixProduct[p.size] / p.prefixProduct[p.size-k]
}
