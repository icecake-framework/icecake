

function ickError(msg) { console.error(msg) }

function ickWarn(msg) { console.warn(msg) }

function tryGet() {
    try {
        var source = arguments[0];
        for (var i = 1; i < arguments.length; i++)
            source = source[arguments[i]];

        return [source, null];
    }
    catch (err) {
        console.warn("tryGet", err.error);
        return [null, err];
    }
}

function trySet() {
    try {
        var source = arguments[0];
        for (var i = 1; i < arguments.length; i++)
            source = source[arguments[i]];

        return [null];
    }
    catch (err) {
        console.warn("trySet", err.error);
        return [err];
    }
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

