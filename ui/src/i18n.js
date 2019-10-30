import i18n from "i18next";
import Backend from 'i18next-xhr-backend';
import LanguageDetector from 'i18next-browser-languagedetector';
import { initReactI18next } from "react-i18next";

i18n
  .use(initReactI18next)
  .use(Backend)
  .use(LanguageDetector)
  .init({
    keySeparator: '.',
    interpolation: {
      escapeValue: false
    },
    backend: {
      loadPath: 'lang/{{lng}}.json'
    },
    fallbackLng: 'en'
  });

export default i18n;