"use-strict";

const header = document.getElementsByClassName("header");
const headerNav = document.getElementsByClassName("header__nav");

/**
 * @param {PointerEvent} e
 */
function activeHeaderClick(e) {
  if (e.composedPath().includes(header.item(0))) {
    return;
  }
  toggleHeader();
  document.removeEventListener("pointerdown", activeHeaderClick);
}

function toggleHeader() {
  const expanded = headerNav.item(0)?.toggleAttribute("aria-expanded");
  if (expanded) {
    document.addEventListener("pointerdown", activeHeaderClick);
  }
}
