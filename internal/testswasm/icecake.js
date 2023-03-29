
function ickError(msg) { console.error(msg) }

function ickWarn(msg) { console.warn(msg) }

/******************************************************************************
 * Local Storage
 */

function ickLocalStorage() {
    var ls;
    try {
        ls = window.localStorage;
    } catch (e) {
        console.log("ick", e.error);
        return null;
    }
    return ls;
}

function ickSessionStorage() {
    var ls;
    try {
        ls = window.sessionStorage;
    } catch (e) {
        console.log("ick", e.error);
        return null;
    }
    return ls;
}

function ickStorageSetItem(storage, key, value) {
    try {
        storage.setItem(key, value);
    } catch (e) {
        if (isQuotaExceeded(e)) {
            // Storage full, maybe notify user or do some clean-up
            return "setItem fails: storage if full"
        }
        return "setItem fails: ?"
    }
    return null
}

function isQuotaExceeded(e) {
    var quotaExceeded = false;
    if (e) {
        if (e.code) {
            switch (e.code) {
                case 22:
                    quotaExceeded = true;
                    break;
                case 1014:
                    // Firefox
                    if (e.name === 'NS_ERROR_DOM_QUOTA_REACHED') {
                        quotaExceeded = true;
                    }
                    break;
            }
        } else if (e.number === -2147024882) {
            // Internet Explorer 8
            quotaExceeded = true;
        }
    }
    return quotaExceeded;
}

