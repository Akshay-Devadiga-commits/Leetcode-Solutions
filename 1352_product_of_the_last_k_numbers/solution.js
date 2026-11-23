class ProductOfNumbers {
    constructor() {
        this.prefixProduct = [1];
        this.size = 0;
    }

    add(num) {
        if (num === 0) {
            this.prefixProduct = [1];
            this.size = 0;
        } else {
            this.prefixProduct.push(this.prefixProduct[this.size] * num);
            this.size++;
        }
    }

    getProduct(k) {
        if (k > this.size) return 0;
        return Math.floor(this.prefixProduct[this.size] / this.prefixProduct[this.size - k]);
    }
}
