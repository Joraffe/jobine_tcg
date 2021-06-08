export const MakeEventBus = () => {
  const _events = {};

  const eventBus = {
    subscribe: (eventName, subscriberId, callback) => {
      if (!_events.hasOwnProperty(eventName)) {
        _events[eventName] = {};
      }

      if (!_events[eventName].hasOwnProperty(subscriberId)) {
        _events[eventName][subscriberId] = [];
      }

      _events[eventName][subscriberId].push(callback);
    },
    unsubscribe: (eventName, subscriberId) => {
      if (!_events.hasOwnProperty(eventName)) {
        console.warn(`No such event ${eventName} exists`);
        return;
      }

      if (!_events[eventName].hasOwnProperty(subscriberId)) {
        console.warn(`${subscriberId} is not currently listening to ${eventName}`);
        return;
      }

      delete _events[eventName][subscriberId];
    },
    publish: (eventName, eventData) => {
      const eventSubscribers = _events[eventName];
      for (const subscriberId in eventSubscribers) {
        let subsciberEvents = eventSubscribers[subscriberId];
        subsciberEvents.forEach((callback) => {
          callback(eventData);
        });
      }
    },
  };

  return eventBus;
}
