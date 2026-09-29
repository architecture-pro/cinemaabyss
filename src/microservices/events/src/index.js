const express = require('express');
const { Kafka } = require('kafkajs');

const app = express();

const PORT = process.env.PORT || 8082;
const KAFKA_BROKERS = (process.env.KAFKA_BROKERS || 'localhost:9092').split(
	',',
);

app.use(express.json());

const kafka = new Kafka({
	clientId: 'cinemaabyss-events-service',
	brokers: KAFKA_BROKERS,
});

const producer = kafka.producer();

const consumer = kafka.consumer({
	groupId: 'cinemaabyss-events-group',
});

async function publishEvent(topic, event) {
	await producer.send({
		topic,
		messages: [
			{
				value: JSON.stringify(event),
			},
		],
	});

	console.log(`Event published to ${topic}:`, event);
}

app.get('/api/events/health', (req, res) => {
	res.json({
		status: true,
	});
});

app.post('/api/events/movie', async (req, res) => {
	try {
		await publishEvent('movie-events', req.body);

		res.status(201).json({
			status: 'success',
		});
	} catch (error) {
		console.error('Failed to publish movie event:', error);

		res.status(500).json({
			status: 'error',
		});
	}
});

app.post('/api/events/user', async (req, res) => {
	try {
		await publishEvent('user-events', req.body);

		res.status(201).json({
			status: 'success',
		});
	} catch (error) {
		console.error('Failed to publish user event:', error);

		res.status(500).json({
			status: 'error',
		});
	}
});

app.post('/api/events/payment', async (req, res) => {
	try {
		await publishEvent('payment-events', req.body);

		res.status(201).json({
			status: 'success',
		});
	} catch (error) {
		console.error('Failed to publish payment event:', error);

		res.status(500).json({
			status: 'error',
		});
	}
});

async function start() {
	try {
		await producer.connect();
		console.log(`Producer connected to Kafka: ${KAFKA_BROKERS.join(', ')}`);

		await consumer.connect();
		console.log(`Consumer connected to Kafka: ${KAFKA_BROKERS.join(', ')}`);

		await consumer.subscribe({
			topics: ['movie-events', 'user-events', 'payment-events'],
			fromBeginning: false,
		});

		await consumer.run({
			eachMessage: async ({ topic, partition, message }) => {
				const event = JSON.parse(message.value.toString());

				console.log(`Event consumed from ${topic}:`, event);
			},
		});

		app.listen(PORT, () => {
			console.log(`Events service started on port ${PORT}`);
		});
	} catch (error) {
		console.error('Failed to start events service:', error);
		process.exit(1);
	}
}

start();
