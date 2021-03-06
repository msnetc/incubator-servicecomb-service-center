/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package broker

import (
	"github.com/apache/incubator-servicecomb-service-center/server/core/backend"
)

var (
	PARTICIPANT  backend.StoreType
	VERSION      backend.StoreType
	PACT         backend.StoreType
	PACT_VERSION backend.StoreType
	PACT_TAG     backend.StoreType
	VERIFICATION backend.StoreType
	PACT_LATEST  backend.StoreType
)

var brokerKvStore = &BKvStore{}

func init() {
	PARTICIPANT = backend.Store().MustInstall(backend.NewExtension("PARTICIPANT", backend.Configure().WithPrefix(GetBrokerParticipantKey(""))))
	VERSION = backend.Store().MustInstall(backend.NewExtension("VERSION", backend.Configure().WithPrefix(GetBrokerVersionKey(""))))
	PACT = backend.Store().MustInstall(backend.NewExtension("PACT", backend.Configure().WithPrefix(GetBrokerPactKey(""))))
	PACT_VERSION = backend.Store().MustInstall(backend.NewExtension("PACT_VERSION", backend.Configure().WithPrefix(GetBrokerPactVersionKey(""))))
	PACT_TAG = backend.Store().MustInstall(backend.NewExtension("PACT_TAG", backend.Configure().WithPrefix(GetBrokerTagKey(""))))
	VERIFICATION = backend.Store().MustInstall(backend.NewExtension("VERIFICATION", backend.Configure().WithPrefix(GetBrokerVerificationKey(""))))
	PACT_LATEST = backend.Store().MustInstall(backend.NewExtension("PACT_LATEST", backend.Configure().WithPrefix(GetBrokerLatestKey(""))))

}

type BKvStore struct {
}

func (s *BKvStore) Participant() backend.Indexer {
	return backend.Store().Entities(PARTICIPANT)
}

func (s *BKvStore) Version() backend.Indexer {
	return backend.Store().Entities(VERSION)
}

func (s *BKvStore) Pact() backend.Indexer {
	return backend.Store().Entities(PACT)
}

func (s *BKvStore) PactVersion() backend.Indexer {
	return backend.Store().Entities(PACT_VERSION)
}

func (s *BKvStore) PactTag() backend.Indexer {
	return backend.Store().Entities(PACT_TAG)
}

func (s *BKvStore) Verification() backend.Indexer {
	return backend.Store().Entities(VERIFICATION)
}

func (s *BKvStore) PactLatest() backend.Indexer {
	return backend.Store().Entities(PACT_LATEST)
}

func Store() *BKvStore {
	return brokerKvStore
}
