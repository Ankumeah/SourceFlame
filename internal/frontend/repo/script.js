async function main() {
  const repo = window.location.pathname.slice(3, -1);

  document.getElementById("repo-name").innerText = repo;

  document.title = repo

  setClone(repo);
  if (await loadRefs(repo)) { return }

  loadReadme(repo);
  loadLog(repo);
}

function setClone(repo) {
  document.getElementById("clone-link").innerText = location.origin + "/api/v1/r/" + repo;
}

async function loadReadme(repo) {

  let ref = document.getElementById("ref-selector").value;
  if (!ref) {
    document.getElementById("readme").innerText = "Select a valid ref";
    return
  }

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
  if (res.status === 404 || res.status === 400) {
    setInvalidRepo();
    return true;
  }
  if (!res.ok) {
    throw new Error(`Error loading refs: ${res.status}`);
  }

  let refs = {};

  try {
    refs = await res.json();
  } catch (err) {
    console.log(`Invalid json from refs api: ${err}`);
  }

  if (refs.refs.length < 1) {
    refs.refs = [{ name: "None", hash: "" }];
  }

  let opts = "";
  refs.refs.forEach(function(ref) {
    opts += `<option value="${ref.hash}">${ref.name}</option>`;
  });

  document.getElementById("ref-selector").innerHTML = opts;

  return false;
}

async function loadLog(repo) {
  const res = await fetch(
    `${location.origin}/api/v1/r/` +
    `${repo}/git-log?ref=` +
    `${document.getElementById("ref-selector").value}`
  );
  if (res === 404) {
    document.getElementById("commits").innerHTML = "No log";
    return
  } else if (!res.ok) {
    throw new Error(`Error loading log: ${res.status}`);
  }

  document.getElementById("log").innerText = await res.text();
}

function setInvalidRepo() {
  document.title = "Invalid repo";
  document.getElementById("repo-name").innerText = "Invalid repo";
  document.getElementById("info").hidden = true;
}

window.addEventListener(
  "DOMContentLoaded",
  main
);
