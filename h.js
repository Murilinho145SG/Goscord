function pegarTextoSemBlink() {
  const phrase = document.getElementById("phrase").cloneNode(true);
  phrase.querySelectorAll(".blink").forEach(b => b.remove());
  return phrase.textContent.replace(/\s+/g, " ").trim();
}

function simularTecla(char) {
  const eventDown = new KeyboardEvent("keydown", { key: char, bubbles: true });
  const eventPress = new KeyboardEvent("keypress", { key: char, bubbles: true });
  const eventUp = new KeyboardEvent("keyup", { key: char, bubbles: true });

  document.dispatchEvent(eventDown);
  document.dispatchEvent(eventPress);
  document.dispatchEvent(eventUp);
}

function digitarAutomatico(texto, delay = 300) {
  let i = 0;

  function digitar() {
    if (i < texto.length) {
      const char = texto[i];
      simularTecla(char);
      i++;
      setTimeout(digitar, delay);
    }
  }

  digitar();
}

// --- Uso ---
const texto = pegarTextoSemBlink();
digitarAutomatico(texto, 300);