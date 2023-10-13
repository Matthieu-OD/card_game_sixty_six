import 'vite/modulepreload-polyfill'
import htmx from 'htmx.org'
import Alpine from 'alpinejs'

import '../css/style.css'

import copyToClipboard from './alpine/data/copy_to_clipboard'


window.Alpine = Alpine
window.htmx = htmx


Alpine.data('copyToClipboard', copyToClipboard)

Alpine.start()
