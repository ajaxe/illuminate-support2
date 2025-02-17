function renderReleaseList() {
  $.get({
    url: "/web/content/releases.md",
    data: { _: new Date().getTime() },
    dataType: "text",
  }).then(function (data) {
    showReleaseList(data);
  });
}

function showReleaseList(data) {
  const $releases = $("#releases .content").first();
  var reader = new commonmark.Parser();
  var writer = new commonmark.HtmlRenderer();
  var parsed = reader.parse(data);

  $releases.html(writer.render(parsed));
}

function renderFeatureList(features, $embedCheckIcon) {
  const $featureList = $("#feature-list");
  features.forEach((v) => {
    $embedCheckIcon.clone().appendTo($featureList);
    $("<span>").text(v).appendTo($featureList);
    $("<br>").appendTo($featureList);
  });
}

const features = [
  "Customize color list for text highlighting.",
  "Disable extension on a single page or entire domain.",
  "Clear all text highlights with single click.",
  "See the summary of selections on the 'Options' page.",
];

const embedCheckIcon =
  "<svg width='1.5em' height='1.5em' viewBox='0 0 16 16' class='bi bi-check2 text-success' fill='currentColor' xmlns='http://www.w3.org/2000/svg' > <path   fill-rule='evenodd'   d='M13.854 3.646a.5.5 0 0 1 0 .708l-7 7a.5.5 0 0 1-.708 0l-3.5-3.5a.5.5 0 1 1 .708-.708L6.5 10.293l6.646-6.647a.5.5 0 0 1 .708 0z' /> </svg>";
