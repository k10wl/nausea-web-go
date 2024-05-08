"use-strict";

const header = document.getElementsByClassName("header");
const headerNav = document.getElementsByClassName("header__nav");

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

function toggleHeader() {
  const expanded = headerNav.item(0)?.toggleAttribute("aria-expanded");
  if (expanded) {
    onOutsidePointerDown(header.item(0), toggleHeader);
  }
}
