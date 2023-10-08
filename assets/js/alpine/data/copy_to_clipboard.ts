export default () => ({
	copy() {
		navigator.clipboard.writeText(this.$refs.linkToCopy.innerHTML)
	}
})
