class MinHeap {
    constructor(compare = (a, b) => a - b) {
        this.heap = [];
        this.compare = compare;
    }
    size() {
        return this.heap.length;
    }
    peek() {
        return this.heap[0];
    }
    push(val) {
        this.heap.push(val);
        this._up(this.heap.length - 1);
    }
    pop() {
        if (this.size() === 0) return null;
        const top = this.heap[0];
        const bottom = this.heap.pop();
        if (this.size() > 0) {
            this.heap[0] = bottom;
            this._down(0);
        }
        return top;
    }
    _up(i) {
        while (i > 0) {
            const p = (i - 1) >> 1;
            if (this.compare(this.heap[i], this.heap[p]) < 0) {
                [this.heap[i], this.heap[p]] = [this.heap[p], this.heap[i]];
                i = p;
            } else {
                break;
            }
        }
    }
    _down(i) {
        const len = this.heap.length;
        while ((i << 1) + 1 < len) {
            let left = (i << 1) + 1;
            let right = left + 1;
            let best = left;
            if (right < len && this.compare(this.heap[right], this.heap[left]) < 0) {
                best = right;
            }
            if (this.compare(this.heap[best], this.heap[i]) < 0) {
                [this.heap[i], this.heap[best]] = [this.heap[best], this.heap[i]];
                i = best;
            } else {
                break;
            }
        }
    }
}

/**
 * @param {number} k
 * @param {number[]} arrival
 * @param {number[]} load
 * @return {number[]}
 */
var busiestServers = function(k, arrival, load) {
    const count = new Array(k).fill(0);
    const busy = new MinHeap((a, b) => a[0] - b[0]);
    const after = new MinHeap((a, b) => a - b);
    const before = new MinHeap((a, b) => a - b);

    for (let i = 0; i < k; i++) {
        after.push(i);
    }

    for (let i = 0; i < arrival.length; i++) {
        const currentTime = arrival[i];
        const requestLoad = load[i];
        const target = i % k;

        if (i > 0 && target === 0) {
            while (before.size() > 0) {
                after.push(before.pop());
            }
        } else {
            while (after.size() > 0 &&