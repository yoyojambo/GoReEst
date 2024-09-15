/**
 * @license
 * Copyright 2019 Google LLC. All Rights Reserved.
 * SPDX-License-Identifier: Apache-2.0
 */
let mapsInitialized = {};

function initializeMap(containerId) {
  if (mapsInitialized[containerId]) return; // Prevent re-initialization

  const container = document.getElementById(containerId);
  if (!container) return;

  container.querySelectorAll('.map-container').forEach(container => {
    const lat = parseFloat(container.getAttribute('data-lat'));
    const lng = parseFloat(container.getAttribute('data-lng'));
    const mapId = container.id;

    // Request needed libraries.
    //@ts-ignore
    google.maps.importLibrary("maps").then(({ Map }) => {
      google.maps.importLibrary("marker").then(({ AdvancedMarkerElement }) => {
        // Initialize map with specific coordinates
        const map = new Map(document.getElementById(mapId), {
          zoom: 4,
          center: { lat: lat, lng: lng },
          mapId: "DEMO_MAP_ID",
        });

        // Add a marker to the map
        new AdvancedMarkerElement({
          map: map,
          position: { lat: lat, lng: lng },
          title: `Map at ${lat}, ${lng}`,
        });
      });
    });
  });

  mapsInitialized[containerId] = true; // Mark as initialized
}

// maps.js

function toggleDropdown(dropdownId) {
  console.log('toggleDropdown function called');
  const dropdown = document.getElementById(dropdownId);
  const isOpen = dropdown.style.display === 'block';
 
  // Hide all dropdowns
  document.querySelectorAll('.dropdown-content').forEach(content => content.style.display = 'none');
  if (!isOpen) {
      dropdown.style.display = 'block';
      initializeMap(dropdownId); // Initialize maps within the dropdown
  }
}


