import VueI18n from 'vue-i18n'
import Vue from 'vue'
Vue.use(VueI18n)

const messages = {
    zh: {
        message: require("./i18n/zh.json")
    },
    en: {
        message: require("./i18n/en.json")
    }
}

export const langs = {
    "en": "English",
    "zh": "中文"
}

export const i18n = new VueI18n({
    locale: "en",
    messages,
    // silentTranslationWarn: true
})