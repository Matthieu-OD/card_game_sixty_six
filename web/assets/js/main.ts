import 'vite/modulepreload-polyfill'
import htmx from 'htmx.org'
import Alpine from 'alpinejs'

import '../css/style.css'

import copyToClipboard from './alpine/data/copy_to_clipboard'


window.Alpine = Alpine
window.htmx = htmx


Alpine.data('copyToClipboard', copyToClipboard)

Alpine.start()

// Dynamically import the sse.js extension when needed
import('htmx.org/dist/ext/sse.js').then(() => {
	console.log('Loaded sse.js');
}).catch((error) => {
	console.error('Error loading sse.js:', error);
});
