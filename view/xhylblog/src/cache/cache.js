import storage from 'good-storage'

class cache {
  constructor(key, session) {
    this.key = key;
    if (session) {
      this.storage = storage.session;
    } else {
      this.storage = storage;
    }
  }

  get(defaultValue = null) {
    return this.storage.get(this.key, defaultValue);
  }

  set(value) {
    this.storage.set(this.key, value);}
}

let prefix = (window.version ? window.version + "_" : '') +"self_";
const USER_CACHE_KEY = prefix + 'user';
export const USER_CACHE = new cache(USER_CACHE_KEY);