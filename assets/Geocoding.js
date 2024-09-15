let map;
let marker;
let geocoder;
let responseDiv;

function initMap() {
    map = new google.maps.Map(document.getElementById("map"), {
        zoom: 8,
        center: { lat: -34.397, lng: 150.644 },
        mapTypeControl: false,
    });
    geocoder = new google.maps.Geocoder();

    const submitButton = document.getElementById("submit");
    const inputText = document.getElementById("address");
    responseDiv = document.getElementById("response-container");

    marker = new google.maps.Marker({
        map,
    });
    map.addListener("click", (e) => {
        geocode({ location: e.latLng });
    });
    submitButton.addEventListener("click", () =>
        geocode({ address: inputText.value })
    );
}

function geocode(request) {
  marker.setMap(null);
  geocoder
      .geocode(request)
      .then((result) => {
          const { results } = result;
          if (results.length > 0) {
              const location = results[0].geometry.location;
              map.setCenter(location);
	      map.setZoom(15)
              marker.setPosition(location);
              marker.setMap(map);
              const lat = location.lat();
              const lng = location.lng();
              responseDiv.innerText = `Latitude: ${lat}, Longitude: ${lng}`;
          } else {
              responseDiv.innerText = 'No results found.';
          }
      })
      .catch((e) => {
          alert("Geocode was not successful for the following reason: " + e);
      });
}


window.initMap = initMap;
