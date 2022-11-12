import { LitElement, html } from 'https://cdn.jsdelivr.net/gh/lit/dist@2/core/lit-core.min.js';

export default class CodeBlock extends LitElement {
    static get properties() {
        return {
            lang: {type: String},
            src: {type: String}
        };
    }
    firstUpdated() {
        fetch(this.src).then(response => this.rawSource = response.text());
    }

    render() {
        return html`
        <code class="language-${this.lang} hljs">
            ${this.rawSource}
        </code>
        `;
    }
}

customElements.define("code-block", CodeBlock)
