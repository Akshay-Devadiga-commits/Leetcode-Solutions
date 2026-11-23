class ProductOfNumbers {
    // Stores cumulative product of the stream
    constructor() {
        // Initialize the product list with 1 to handle multiplication logic
        this.prefixProduct = [1];
        this.size = 0;
    }

    add(num) {
        if (num === 0) {
            // If num is 0, reset the cumulative products since multiplication
            // with 0 invalidates previous products
            this.prefixProduct = [1];
            this.size = 0;
        } else {
            // Append the cumulative product of the current number with the last
            // product
            this.prefixProduct.push(this.prefixProduct[this.size] * num);
            this.size++;
        }
    }

    getProduct(k) {
        // Check if the requested product length exceeds the size of the valid
        // product list
        if (k > this.size) return 0;

        // Compute the product of the last k elements using division
        return Math.floor(this.prefixProduct[this.size] / this.prefixProduct[this.size - k]);
    }
}