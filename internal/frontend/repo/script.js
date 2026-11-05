async function main() {
  const repo = window.location.pathname.slice(3, -1);

  document.getElementById("repo-name").innerText = repo;

  setClone(repo);
  await loadRefs(repo);
  await loadReadme(repo);
}

function setClone(repo) {
  document.getElementById("clone-link").innerText = location.origin + "/api/v1/r/" + repo;
}

async function loadReadme(repo) {
  let text = ""

  const res = await fetch(
    `${location.origin}/api/v1/r/` +
    `${repo}/blob/README.md?ref=` +
    `${document.getElementById("ref-selector").value}`,
  );
  if (res.status === 404) {
    text = "No README.md at repo root";
  } else if (!res.ok) {
    text = "Error while loading README"
    throw new Error(`Error loading refs: ${res.status}`);
  } else {
    text = await res.text();
  }

  document.getElementById("readme").innerText = text;
}

async function loadRefs(repo) {
  const res = await fetch(`${location.origin}/api/v1/r/${repo}/refs`);
  if (!res.ok) {
    throw new Error(`Error loading refs: ${res.status}`);
  }

  let refs = {};

  try {
    refs = await res.json();
  } catch (err) {
    console.log(`Invalid json from refs api: ${err}`);
  }

  let opts = "";
  refs.refs.forEach(function(ref) {
    opts += `<option value="${ref.hash}">${ref.name}</option>`;
  });

  document.getElementById("ref-selector").innerHTML = opts;
}

window.addEventListener(
  "DOMContentLoaded",
  () => main()
);
