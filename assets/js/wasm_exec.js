"use strict";
(() => {
    const l = () => {
        var e = new Error("not implemented");
        return e.code = "ENOSYS", e
    };
    if (!globalThis.fs) {
        let i = "";
        globalThis.fs = {
            constants: {O_WRONLY: -1, O_RDWR: -1, O_CREAT: -1, O_TRUNC: -1, O_APPEND: -1, O_EXCL: -1},
            writeSync(e, t) {
                var s = (i += c.decode(t)).lastIndexOf("\n");
                return -1 != s && (console.log(i.substring(0, s)), i = i.substring(s + 1)), t.length
            },
            write(e, t, s, i, n, r) {
                0 !== s || i !== t.length || null !== n ? r(l()) : r(null, this.writeSync(e, t))
            },
            chmod(e, t, s) {
                s(l())
            },
            chown(e, t, s, i) {
                i(l())
            },
            close(e, t) {
                t(l())
            },
            fchmod(e, t, s) {
                s(l())
            },
            fchown(e, t, s, i) {
                i(l())
            },
            fstat(e, t) {
                t(l())
            },
            fsync(e, t) {
                t(null)
            },
            ftruncate(e, t, s) {
                s(l())
            },
            lchown(e, t, s, i) {
                i(l())
            },
            link(e, t, s) {
                s(l())
            },
            lstat(e, t) {
                t(l())
            },
            mkdir(e, t, s) {
                s(l())
            },
            open(e, t, s, i) {
                i(l())
            },
            read(e, t, s, i, n, r) {
                r(l())
            },
            readdir(e, t) {
                t(l())
            },
            readlink(e, t) {
                t(l())
            },
            rename(e, t, s) {
                s(l())
            },
            rmdir(e, t) {
                t(l())
            },
            stat(e, t) {
                t(l())
            },
            symlink(e, t, s) {
                s(l())
            },
            truncate(e, t, s) {
                s(l())
            },
            unlink(e, t) {
                t(l())
            },
            utimes(e, t, s, i) {
                i(l())
            }
        }
    }
    if (globalThis.process || (globalThis.process = {
        getuid() {
            return -1
        }, getgid() {
            return -1
        }, geteuid() {
            return -1
        }, getegid() {
            return -1
        }, getgroups() {
            throw l()
        }, pid: -1, ppid: -1, umask() {
            throw l()
        }, cwd() {
            throw l()
        }, chdir() {
            throw l()
        }
    }), !globalThis.crypto) throw new Error("globalThis.crypto is not available, polyfill required (crypto.getRandomValues only)");
    if (!globalThis.performance) throw new Error("globalThis.performance is not available, polyfill required (performance.now only)");
    if (!globalThis.TextEncoder) throw new Error("globalThis.TextEncoder is not available, polyfill required");
    if (!globalThis.TextDecoder) throw new Error("globalThis.TextDecoder is not available, polyfill required");
    const m = new TextEncoder("utf-8"), c = new TextDecoder("utf-8");
    globalThis.Go = class {
        constructor() {
            this.argv = ["js"], this.env = {}, this.exit = e => {
                0 !== e && console.warn("exit code:", e)
            }, this._exitPromise = new Promise(e => {
                this._resolveExitPromise = e
            }), this._pendingEvent = null, this._scheduledTimeouts = new Map, this._nextCallbackTimeoutID = 1;
            const i = (e, t) => {
                this.mem.setUint32(e + 0, t, !0), this.mem.setUint32(e + 4, Math.floor(t / 4294967296), !0)
            };
            const n = e => {
                return this.mem.getUint32(e + 0, !0) + 4294967296 * this.mem.getInt32(e + 4, !0)
            }, r = e => {
                var t = this.mem.getFloat64(e, !0);
                if (0 !== t) return isNaN(t) ? (e = this.mem.getUint32(e, !0), this._values[e]) : t
            }, l = (s, i) => {
                var n = 2146959360;
                if ("number" == typeof i && 0 !== i) return isNaN(i) ? (this.mem.setUint32(s + 4, n, !0), void this.mem.setUint32(s, 0, !0)) : void this.mem.setFloat64(s, i, !0);
                if (void 0 === i) this.mem.setFloat64(s, 0, !0); else {
                    let e = this._ids.get(i),
                        t = (void 0 === e && (void 0 === (e = this._idPool.pop()) && (e = this._values.length), this._values[e] = i, this._goRefCounts[e] = 0, this._ids.set(i, e)), this._goRefCounts[e]++, 0);
                    switch (typeof i) {
                        case"object":
                            null !== i && (t = 1);
                            break;
                        case"string":
                            t = 2;
                            break;
                        case"symbol":
                            t = 3;
                            break;
                        case"function":
                            t = 4
                    }
                    this.mem.setUint32(s + 4, n | t, !0), this.mem.setUint32(s, e, !0)
                }
            }, o = e => {
                var t = n(e + 0), e = n(e + 8);
                return new Uint8Array(this._inst.exports.mem.buffer, t, e)
            }, a = e => {
                var t = n(e + 0), s = n(e + 8), i = new Array(s);
                for (let e = 0; e < s; e++) i[e] = r(t + 8 * e);
                return i
            }, h = e => {
                var t = n(e + 0), e = n(e + 8);
                return c.decode(new DataView(this._inst.exports.mem.buffer, t, e))
            }, t = Date.now() - performance.now();
            this.importObject = {
                _gotest: {add: (e, t) => e + t}, gojs: {
                    "runtime.wasmExit": e => {
                        e = this.mem.getInt32((e >>>= 0) + 8, !0);
                        this.exited = !0, delete this._inst, delete this._values, delete this._goRefCounts, delete this._ids, delete this._idPool, this.exit(e)
                    }, "runtime.wasmWrite": e => {
                        e >>>= 0;
                        var t = n(e + 8), s = n(e + 16), e = this.mem.getInt32(e + 24, !0);
                        fs.writeSync(t, new Uint8Array(this._inst.exports.mem.buffer, s, e))
                    }, "runtime.resetMemoryDataView": e => {
                        this.mem = new DataView(this._inst.exports.mem.buffer)
                    }, "runtime.nanotime1": e => {
                        e >>>= 0, i(e + 8, 1e6 * (t + performance.now()))
                    }, "runtime.walltime": e => {
                        e >>>= 0;
                        var t = (new Date).getTime();
                        i(e + 8, t / 1e3), this.mem.setInt32(e + 16, t % 1e3 * 1e6, !0)
                    }, "runtime.scheduleTimeoutEvent": e => {
                        e >>>= 0;
                        const t = this._nextCallbackTimeoutID;
                        this._nextCallbackTimeoutID++, this._scheduledTimeouts.set(t, setTimeout(() => {
                            for (this._resume(); this._scheduledTimeouts.has(t);) console.warn("scheduleTimeoutEvent: missed timeout event"), this._resume()
                        }, n(e + 8))), this.mem.setInt32(e + 16, t, !0)
                    }, "runtime.clearTimeoutEvent": e => {
                        e = this.mem.getInt32((e >>>= 0) + 8, !0);
                        clearTimeout(this._scheduledTimeouts.get(e)), this._scheduledTimeouts.delete(e)
                    }, "runtime.getRandomData": e => {
                        e >>>= 0, crypto.getRandomValues(o(e + 8))
                    }, "syscall/js.finalizeRef": e => {
                        var t, e = this.mem.getUint32((e >>>= 0) + 8, !0);
                        this._goRefCounts[e]--, 0 === this._goRefCounts[e] && (t = this._values[e], this._values[e] = null, this._ids.delete(t), this._idPool.push(e))
                    }, "syscall/js.stringVal": e => {
                        e >>>= 0, l(e + 24, h(e + 8))
                    }, "syscall/js.valueGet": e => {
                        e >>>= 0;
                        var t = Reflect.get(r(e + 8), h(e + 16));
                        e = this._inst.exports.getsp() >>> 0, l(e + 32, t)
                    }, "syscall/js.valueSet": e => {
                        e >>>= 0, Reflect.set(r(e + 8), h(e + 16), r(e + 32))
                    }, "syscall/js.valueDelete": e => {
                        e >>>= 0, Reflect.deleteProperty(r(e + 8), h(e + 16))
                    }, "syscall/js.valueIndex": e => {
                        e >>>= 0, l(e + 24, Reflect.get(r(e + 8), n(e + 16)))
                    }, "syscall/js.valueSetIndex": e => {
                        e >>>= 0, Reflect.set(r(e + 8), n(e + 16), r(e + 24))
                    }, "syscall/js.valueCall": t => {
                        t >>>= 0;
                        try {
                            var e = r(t + 8), s = Reflect.get(e, h(t + 16)), i = a(t + 32), n = Reflect.apply(s, e, i);
                            t = this._inst.exports.getsp() >>> 0, l(t + 56, n), this.mem.setUint8(t + 64, 1)
                        } catch (e) {
                            t = this._inst.exports.getsp() >>> 0, l(t + 56, e), this.mem.setUint8(t + 64, 0)
                        }
                    }, "syscall/js.valueInvoke": t => {
                        t >>>= 0;
                        try {
                            var e = r(t + 8), s = a(t + 16), i = Reflect.apply(e, void 0, s);
                            t = this._inst.exports.getsp() >>> 0, l(t + 40, i), this.mem.setUint8(t + 48, 1)
                        } catch (e) {
                            t = this._inst.exports.getsp() >>> 0, l(t + 40, e), this.mem.setUint8(t + 48, 0)
                        }
                    }, "syscall/js.valueNew": t => {
                        t >>>= 0;
                        try {
                            var e = r(t + 8), s = a(t + 16), i = Reflect.construct(e, s);
                            t = this._inst.exports.getsp() >>> 0, l(t + 40, i), this.mem.setUint8(t + 48, 1)
                        } catch (e) {
                            t = this._inst.exports.getsp() >>> 0, l(t + 40, e), this.mem.setUint8(t + 48, 0)
                        }
                    }, "syscall/js.valueLength": e => {
                        e >>>= 0, i(e + 16, parseInt(r(e + 8).length))
                    }, "syscall/js.valuePrepareString": e => {
                        e >>>= 0;
                        var t = m.encode(String(r(e + 8)));
                        l(e + 16, t), i(e + 24, t.length)
                    }, "syscall/js.valueLoadString": e => {
                        e >>>= 0;
                        var t = r(e + 8);
                        o(e + 16).set(t)
                    }, "syscall/js.valueInstanceOf": e => {
                        this.mem.setUint8((e >>>= 0) + 24, r(e + 8) instanceof r(e + 16) ? 1 : 0)
                    }, "syscall/js.copyBytesToGo": e => {
                        e >>>= 0;
                        var t = o(e + 8), s = r(e + 32);
                        s instanceof Uint8Array || s instanceof Uint8ClampedArray ? (s = s.subarray(0, t.length), t.set(s), i(e + 40, s.length), this.mem.setUint8(e + 48, 1)) : this.mem.setUint8(e + 48, 0)
                    }, "syscall/js.copyBytesToJS": e => {
                        e >>>= 0;
                        var t = r(e + 8), s = o(e + 16);
                        t instanceof Uint8Array || t instanceof Uint8ClampedArray ? (s = s.subarray(0, t.length), t.set(s), i(e + 40, s.length), this.mem.setUint8(e + 48, 1)) : this.mem.setUint8(e + 48, 0)
                    }, debug: e => {
                        console.log(e)
                    }
                }
            }
        }

        async run(e) {
            if (!(e instanceof WebAssembly.Instance)) throw new Error("Go.run: WebAssembly.Instance expected");
            this._inst = e, this.mem = new DataView(this._inst.exports.mem.buffer), this._values = [NaN, 0, null, !0, !1, globalThis, this], this._goRefCounts = new Array(this._values.length).fill(1 / 0), this._ids = new Map([[0, 1], [null, 2], [!0, 3], [!1, 4], [globalThis, 5], [this, 6]]), this._idPool = [], this.exited = !1;
            let s = 4096;
            const t = e => {
                var t = s, e = m.encode(e + "\0");
                return new Uint8Array(this.mem.buffer, s, e.length).set(e), (s += e.length) % 8 != 0 && (s += 8 - s % 8), t
            };
            e = this.argv.length;
            const i = [];
            this.argv.forEach(e => {
                i.push(t(e))
            }), i.push(0);
            Object.keys(this.env).sort().forEach(e => {
                i.push(t(e + "=" + this.env[e]))
            }), i.push(0);
            var n = s;
            i.forEach(e => {
                this.mem.setUint32(s, e, !0), this.mem.setUint32(s + 4, 0, !0), s += 8
            });
            if (12288 <= s) throw new Error("total length of command line and environment variables exceeds limit");
            this._inst.exports.run(e, n), this.exited && this._resolveExitPromise(), await this._exitPromise
        }

        _resume() {
            if (this.exited) throw new Error("Go program has already exited");
            this._inst.exports.resume(), this.exited && this._resolveExitPromise()
        }

        _makeFuncWrapper(t) {
            const s = this;
            return function () {
                var e = {id: t, this: this, args: arguments};
                return s._pendingEvent = e, s._resume(), e.result
            }
        }
    }
})();