async function main() {
  const repo = window.location.pathname.slice(3, -5);
  document.title = repo + " PRs"

  const url = new URL(window.location.href);
  if (url.searchParams.get("page") === null) {
    url.searchParams.set("page", "1");
    window.location.href = url;
  }

  const page = url.searchParams.get("page");

  loadPRs(repo, page);
  setNextAndLast(page);
}

async function loadPRs(repo, page) {
  const res = await fetch(
    `${location.origin}/r/` +
    `${repo}/pr` +
    `?page=${page}`
  );
  if (!res.ok) {
    throw new Error(`Error loading PRs: ${res.status}`);
  }

  let prs = {};
  try {
    prs = await res.json();
  } catch (err) {
    console.log(`Invalid json from prs api: ${err}`);
  }

  prs.prs.forEach((pr) => document.getElementById("prs").appendChild(createPRElement(pr)));
}

function setNextAndLast(page) {
  document.getElementById("next").href = "?page=" + (+page + 1);
  document.getElementById("last").href = "?page=" + Math.max(page - 1, 1);
}

// WARNING
// This func is AI generated
function createPRElement(pr) {
  const details = document.createElement("details");
  details.className = "pr";

  const summary = document.createElement("summary");
  const title = document.createElement("h3");
  title.textContent = pr.title || "Untitled PR";
  summary.appendChild(title);
  details.appendChild(summary);

  const content = document.createElement("div");
  content.className = "pr-content";

  for (const [label, value] of [
    ["Owner", pr.owner],
    ["UUID", pr.uuid],
  ]) {
    const row = document.createElement("p");
    const name = document.createElement("strong");
    name.textContent = `${label}: `;
    row.append(name, document.createTextNode(value || "—"));
    content.appendChild(row);
  }

  const bodyLabel = document.createElement("strong");
  bodyLabel.textContent = "Body:";
  content.appendChild(bodyLabel);

  const body = document.createElement("pre");
  body.className = "pr-body";
  body.textContent = pr.body || "—";
  content.appendChild(body);

  if (pr.url) {
    const row = document.createElement("p");
    const label = document.createElement("strong");
    label.textContent = "URL: ";

    const link = document.createElement("a");
    link.href = pr.url;
    link.textContent = pr.url;
    link.target = "_blank";
    link.rel = "noopener noreferrer";

    row.append(label, link);
    content.appendChild(row);
  }

  details.appendChild(content);
  return details;
}

window.addEventListener(
  "DOMContentLoaded",
  main
);
