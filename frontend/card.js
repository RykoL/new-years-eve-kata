
export class WishCard extends HTMLElement {

  constructor() {
    super();
    const shadowRoot = this.attachShadow({ mode: "open" });
  }

  connectedCallback() {
    this.shadowRoot.innerHTML = `
        <style>
            :host {
                display: block;
            }
            article {
                background: rgba(20, 17, 21, 0.6);
                text-align: center;
                color: white;
                border-radius: 5px;
                padding: 1.5rem;
            }
            h3 {
                margin: 0;
            }
        </style>
        <article>
            <h3>${this.getAttribute('message')}</h3>
        </article>
    `
  }

  static get observedAttributes() {
      return ['message'];
  }
}

customElements.define('wish-card', WishCard);
