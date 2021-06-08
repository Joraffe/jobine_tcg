export const MakeTestModel = (modelId, eventBus) => {

  const model = {
    'callback': (eventData) => {
      console.log(`Model ${modelId} Callback called with ${eventData}`);
    },
    'listenTo': (eventName, callback) => {
      eventBus.subscribe(eventName, modelId, callback);
    },
    'stopListening': (eventName) => {
      eventBus.unsubscribe(eventName, modelId);
    }
  };

  return model;
};
