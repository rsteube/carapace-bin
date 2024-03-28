window.addEventListener("load", function () {
  // <img src="./carapace-bin.cast" alt="" />
  for (elem of Array.prototype.slice.call(document.getElementsByTagName("img")).reverse())
    if (elem.src.endsWith(".cast")) {
      const newItem = document.createElement("div");
      newItem.id = elem.src;
      elem.parentNode.replaceChild(newItem, elem);
      AsciinemaPlayer.create(newItem.id, newItem, {cols: 108, rows: 24});
    }
})
