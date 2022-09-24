import { LitElement, html } from 'https://cdn.jsdelivr.net/gh/lit/dist@2/core/lit-core.min.js';

export default class WASMView extends LitElement {
    static get properties() {
        return {
            height: {type: Number},
            width: {type: Number},
            src: {type: String}
        };
    }
    firstUpdated() {
        const doc = this.shadowRoot.getElementById("wasm-iframe").contentWindow.document
        doc.open()
        doc.write(`
        <!DOCTYPE html>
        <script src="./wasm_exec.js"></script>
        <script>
        // Polyfill
        if (!WebAssembly.instantiateStreaming) {
            WebAssembly.instantiateStreaming = async (resp, importObject) => {
                const source = await (await resp).arrayBuffer();
                return await WebAssembly.instantiate(source, importObject);
            };
        }

        const go = new Go();
        WebAssembly.instantiateStreaming(fetch("${this.src}"), go.importObject).then(result => {
            go.run(result.instance);
        });
        </script>
        `)
        doc.close()
    }

    render() {
        return html`
        <iframe id='wasm-iframe' 
            width="${this.width}" 
            height="${this.height}"
            allow="autoplay"
            scrolling="no"
        >
        </iframe>
        `;
    }
}

customElements.define("wasm-view", WASMView)
