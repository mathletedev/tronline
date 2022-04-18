import { client } from "./lib/client";

let canvas = document.querySelector("canvas") as HTMLCanvasElement;
let ctx = canvas.getContext("2d")!;
let ping = 0;
let last = 0;

canvas.width = window.innerWidth;
canvas.height = window.innerHeight;

window.addEventListener("resize", () => {
  canvas.width = window.innerWidth;
  canvas.height = window.innerHeight;
});

const tick = () => {
  requestAnimationFrame(tick);

  ctx.fillStyle = "#000000";
  ctx.fillRect(0, 0, canvas.width, canvas.height);

  ctx.fillStyle = "#ffffff";
  ctx.font = "18px Cascadia Code";
  ctx.fillText(`ping: ${ping}ms`, 10, 24);
};

requestAnimationFrame(tick);

setInterval(() => {
  last = performance.now();
  client.send("ping");
}, 500);

client.onmessage = (ev) => {
  ping = performance.now() - last;
};
