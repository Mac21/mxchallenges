let dx = 0;
let dy = 0;
let i = 0;

let ruleset = [0, 1, 1, 0, 0, 1, 0, 1];
//let ruleset = [0, 1, 1, 0, 0, 1, 0, 1];
let cells = [1, 0, 0, 0, 0, 0, 0, 0];
//let cells = [0, 1, 1, 1, 1, 1, 1, 1];

function setup() {
  createCanvas(400, 400);
}

function draw() {
  if (cells[i] == 0) {
    fill(255);
  } else {
    fill(0);
  }
  square(dx, dy, 50);

  dx += 50;
  i++;

  if (i == cells.length) {
    i = 0;
    cells = newGeneration(cells);
  }

  if (dx == 400) {
    dx = 0;
    dy += 50;
  }

  if (dy == 400) {
    noLoop();
    return
  }
}

// function newRule(left, middle, right) {
  // let temp = `${left}${middle}${right}`;
  // let i = Number.parseInt(temp, 2);
  // return ruleset[i];
// }

function newThreeNeighRule(a, b, c) {
  if (a == 1 && b == 1 && c == 1) return ruleset[0];
  else if (a == 1 && b == 1 && c == 0) return ruleset[1];
  else if (a == 1 && b == 0 && c == 1) return ruleset[2];
  else if (a == 1 && b == 0 && c == 0) return ruleset[3];
  else if (a == 0 && b == 1 && c == 1) return ruleset[4];
  else if (a == 0 && b == 1 && c == 0) return ruleset[5];
  else if (a == 0 && b == 0 && c == 1) return ruleset[6];
  else if (a == 0 && b == 0 && c == 0) return ruleset[7];
  return 0;
}
 
function newLeftTwoNeighRule(a, b) {
  if (a == 1 && b == 1) return 0;
  else if (a == 1 && b == 0) return 1;
  else if (a == 0 && b == 1) return 0;
  else if (a == 0 && b == 0) return 1;
  return 0;
}

function newRightTwoNeighRule(a, b) {
  if (a == 1 && b == 1) return 1;
  else if (a == 1 && b == 0) return 1;
  else if (a == 0 && b == 1) return 1;
  else if (a == 0 && b == 0) return 1;
  return 0;
}

function newGeneration(cells) {
  let newcells = Array(cells.length);
  let left, middle, right;
  for (let i = 0; i < cells.length; i++) {
    if (i == 0) {
      left = cells[i];
      middle = cells[i+1];
      newcells[i] = newLeftTwoNeighRule(left, middle);
    } else if (i == cells.length-1) {
      left = cells[i-1];
      middle = cells[i];
      newcells[i] = newRightTwoNeighRule(left, middle);
    } else {
      left = cells[i-1];
      middle = cells[i];
      right = cells[i+1];
      newcells[i] = newThreeNeighRule(left, middle, right);
    }
  }
  return newcells;
}
