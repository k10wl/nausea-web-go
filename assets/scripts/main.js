"use-strict";

const header = document.getElementsByClassName("header");
const headerNav = document.getElementsByClassName("header__nav");
const MOBILE_VIEW = 575;

/**
 * Invokes callback upon click outside of element
 * @param {HTMLElement} element - target element
 * @param {() => any} callback function to trigger on outside click
 */
function onOutsidePointerDown(element, callback) {
  /**
   * @param {PointerEvent} event
   */
  function onPointerDown(event) {
    if (event.composedPath().includes(element)) {
      return;
    }
    callback();
    document.removeEventListener("pointerdown", onPointerDown);
  }
  document.addEventListener("pointerdown", onPointerDown);
}

function hideHeader() {
  headerNav.item(0)?.removeAttribute("aria-expanded");
}
function toggleHeader() {
  if (window.outerWidth > MOBILE_VIEW) {
    return;
  }
  const expanded = headerNav.item(0)?.toggleAttribute("aria-expanded");
  if (expanded) {
    onOutsidePointerDown(header.item(0), hideHeader);
  }
}

customElements.define(
  "img-with-thumbnail",
  class ImageWithThumbnail extends HTMLElement {
    constructor() {
      super();
      this.attachShadow({ mode: "open" });
      this.div = document.createElement("div");
      this.img = document.createElement("img");
      this.img.loading = "lazy";
      this.img.onload = () => this.img.style.removeProperty("opacity");
      this.div.appendChild(this.img);
      this.shadowRoot.appendChild(this.div);
    }
    connectedCallback() {
      this.div.style.setProperty(
        "background",
        `url("${this.getAttribute("thumbnail")}")`,
      );
      this.div.style.setProperty("height", "100%");
      this.div.style.setProperty("background-size", "contain");
      this.div.style.setProperty("background-repeat", "no-repeat");
      this.img.style.setProperty("opacity", 0);
      this.img.style.setProperty("display", "block");
      this.img.style.setProperty("max-width", "100%");
      this.img.style.setProperty("max-height", "100%");
      this.img.src = this.getAttribute("src");
    }
  },
);
