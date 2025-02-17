class App {
    constructor() {
        self.app = this;
        this._hashSpeed = 0
    }

    async init() {

    }

    // grid cache
    async getGridCache(id) {
        return this.#getBrowserCache(id + "_grid");
    }

    async setGridCache(id, data) {
        return this.#setBrowserCache(id + "_grid", data);
    }

    // filter cache
    async getFilterCache(id) {
        return this.#getBrowserCache(id + "_filter");
    }

    async setFilterCache(id, data) {
        return this.#setBrowserCache(id + "_filter", data);
    }

    // get browser cache
    async #getBrowserCache(key) {
        return localStorage.getItem(await this.#getHashValue(key));
    }

    // set browser cache
    async #setBrowserCache(key, value) {
        localStorage.setItem(await this.#getHashValue(key), value);
    }

    // get session cache
    async #getSessionCache(key) {
        return sessionStorage.getItem(await this.#getHashValue(key));
    }

    // set session cache
    async #setSessionCache(key, value) {
        sessionStorage.setItem(await this.#getHashValue(key), value);
    }

    // get hash value
    async #getHashValue(key) {
        key = key.toString();
        let h1 = 0xdeadbeef ^ this._hashSpeed , h2 = 0x41c6ce57 ^ this._hashSpeed ;
        for (let i = 0, ch; i < key.length; i++) {
            ch = key.charCodeAt(i);
            h1 = Math.imul(h1 ^ ch, 2654435761);
            h2 = Math.imul(h2 ^ ch, 1597334677);
        }
        h1 = Math.imul(h1 ^ (h1 >>> 16), 2246822507);
        h1 ^= Math.imul(h2 ^ (h2 >>> 13), 3266489909);
        h2 = Math.imul(h2 ^ (h2 >>> 16), 2246822507);
        h2 ^= Math.imul(h1 ^ (h1 >>> 13), 3266489909);

        return 4294967296 * (2097151 & h2) + (h1 >>> 0);
    }
}

new App().init();