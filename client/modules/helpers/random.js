/**
 * Helper to generate a unique random id
 */
export const randomId = () => {
  const randBase36 = Math.random().toString(36);
  const randId = randBase36.substr(2, 9);
  return `_${randId}`;
};
