import { writable } from "svelte/store";

const spacesData = [
  {
    name: "Neon City",
    url: "https://monaverse.com/spaces/neon-city-streets/details",
    image: "/neon-city.png",
  },
  {
    name: "Temple Garden",
    url: "https://monaverse.com/spaces/caelestia:-the-lost-fields/details",
    image: "/temple-garden.png",
  },
  {
    name: "Caelestia",
    url: "https://monaverse.com/spaces/caelestia:-the-lost-fields/details",
    image: "/caelestia.png",
  },
];

const createSpacesStore = () => {
  const { subscribe, set } = writable(spacesData);
  const spacesMap = new Map(spacesData.map((space) => [space.name, space]));

  return {
    subscribe,
    getSpace: (name) => spacesMap.get(name),
    setSpaces: (newSpaces) => set(newSpaces),
  };
};

export const spaces = createSpacesStore();
