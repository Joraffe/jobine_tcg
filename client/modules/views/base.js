import _ from 'lodash';

import { randomId } from 'helpers/random';


export const MakeBaseView = (eventBus) => {
  const view = {
    'id': randomId(),
  };

  Object.assign(
    view,
    MakeEventsMixin(view.id, eventBus),
    MakeRenderMixin(),
  );
};


const MakeEventsMixin = (viewId, eventBus) => {
  const eventsMixin = {
    'listenTo': (eventName, callback) => {
      eventBus.subscribe(eventName, viewId, callback);
    },
    'stopListening': (eventName) => {
      eventBus.unsubscribe(eventName, viewId);
    },
  };

  return eventsMixin;
};


const MakeRenderMixin = () => {
  const renderMixin = {
    'render': () => {
      const templateArgs = this.serializedData();
      const compiledTemplate = _.template(this.template());
      return compiledTemplate(templateArgs);
    },
    'serializedData': () => {
      // This is a no-op; to-be overwritten in concrete views
      return {};
    },
    'template': () => {
      // This is a no-op; to be overwritten in concrete views
      return '';
    },
  };

  return renderMixin;
};
