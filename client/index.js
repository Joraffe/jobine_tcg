import { MakeEventBus } from 'events/pubsub';
import { testLog } from 'helpers/logging';
import { MakeTestModel } from 'models/testing';


testLog('Webpack JS loaded!');

// Singleton event bus
const EventBus = MakeEventBus();

const modelA = MakeTestModel('A', EventBus);
const modelB = MakeTestModel('B', EventBus);

// simple test for subscribing to events
modelA.listenTo('testEvent', modelA.callback);
modelB.listenTo('testEvent', modelB.callback);
EventBus.publish('testEvent', 'my-test-event-data');

// simple test for unsubscribing to events
modelA.stopListening('testEvent');
EventBus.publish('testEvent', 'my-other-test-Event-data');
