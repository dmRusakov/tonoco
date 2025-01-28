class User {
    constructor() {
        self.user = this
    }

    init() {
        this.location = async () => {
            return {
                userAgent: navigator.userAgent,
                platform: navigator.platform,
                isSupportWebp: (() => {
                    const elem = document.createElement('canvas');
                    if (!!(elem.getContext && elem.getContext('2d'))) {
                        return elem.toDataURL('image/webp').indexOf('data:image/webp') == 0;
                    }
                    return false;
                })(),
            }
        }
    }
}

new User()
document.addEventListener('DOMContentLoaded', async () => {
    self.user.init();
}, false);
