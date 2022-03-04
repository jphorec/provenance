#!/usr/bin/make -f

################################################
# Simulation tests with State Listening plugins
#
# This file is an extension for sims.mk
################################################

test-sim-nondeterminism-state-listening-file:
	@echo "Running non-determinism-state-listening-file test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminismWithStateListening -Enabled=true \
		-NumBlocks=50 -BlockSize=100 -Commit=true -Period=0 -v -timeout 24h \
		-StateListeningPlugin=file -HaltAppOnDeliveryError=true

test-sim-nondeterminism-state-listening-trace:
	@echo "Running non-determinism-state-listening-trace test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminismWithStateListening -Enabled=true \
		-NumBlocks=50 -BlockSize=100 -Commit=true -Period=0 -v -timeout 24h \
		-StateListeningPlugin=trace -HaltAppOnDeliveryError=true

test-sim-nondeterminism-state-listening-kafka:
	@echo "Running non-determinism-state-listening-kafka test..."
	@echo "Starting Kafka..."
	docker-compose -f networks/local/kafka/docker-compose.yml up -d

	@-go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminismWithStateListening -Enabled=true \
		-NumBlocks=50 -BlockSize=100 -Commit=true -Period=0 -v -timeout 24h \
		-StateListeningPlugin=kafka -HaltAppOnDeliveryError=true

	@echo "Stopping Kafka..."
	docker-compose -f networks/local/kafka/docker-compose.yml down

test-sim-nondeterminism-state-listening-all: \
	test-sim-nondeterminism-state-listening-file \
	test-sim-nondeterminism-state-listening-trace \
	test-sim-nondeterminism-state-listening-kafka

.PHONY: \
test-sim-nondeterminism-state-listening-all \
test-sim-nondeterminism-state-listening-file \
test-sim-nondeterminism-state-listening-trace \
test-sim-nondeterminism-state-listening-kafka
