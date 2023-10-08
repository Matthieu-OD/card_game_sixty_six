import 'vite/modulepreload-polyfill'
import htmx from 'htmx.org'

import '../css/style.css'
import Alpine from 'alpinejs'

window.Alpine = Alpine
window.htmx = htmx

Alpine.start()
