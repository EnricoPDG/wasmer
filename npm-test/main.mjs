import { bindings } from "@enricopdg/tutorial-1";

async function main() {
  const calculator = await bindings.calculator();
  const four = calculator.add(2, 2);
  console.log('2 + 2 =', four);
}

main();
