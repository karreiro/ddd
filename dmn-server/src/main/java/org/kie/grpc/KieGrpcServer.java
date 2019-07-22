/*
 * Copyright 2019 Red Hat, Inc. and/or its affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package org.kie.grpc;

import java.io.IOException;
import java.util.Map;
import java.util.Optional;
import java.util.UUID;
import java.util.logging.Logger;

import io.grpc.Server;
import io.grpc.ServerBuilder;
import io.grpc.stub.StreamObserver;
import org.kie.api.KieServices;
import org.kie.api.builder.ReleaseId;
import org.kie.api.io.Resource;
import org.kie.api.runtime.KieContainer;
import org.kie.dmn.api.core.DMNContext;
import org.kie.dmn.api.core.DMNModel;
import org.kie.dmn.api.core.DMNResult;
import org.kie.dmn.api.core.DMNRuntime;
import org.kie.dmn.core.api.DMNFactory;
import org.kie.dmn.core.compiler.RuntimeTypeCheckOption;
import org.kie.dmn.core.impl.DMNRuntimeImpl;
import org.kie.dmn.core.util.KieHelper;

public class KieGrpcServer {

    private static final Logger logger = Logger.getLogger(KieGrpcServer.class.getName());

    private Server server;

    public static void main(final String[] args) throws IOException, InterruptedException {
        final KieGrpcServer server = new KieGrpcServer();
        server.start();
        server.blockUntilShutdown();
    }

    private static DMNRuntime createRuntime(final String fileName) {

        final KieServices kieServices = KieServices.Factory.get();
        final Resource resource = kieServices.getResources().newClassPathResource(fileName, KieGrpcServer.class);
        final ReleaseId releaseId = getReleaseId(kieServices);
        final KieContainer kieContainer = KieHelper.getKieContainer(releaseId, resource);

        return typeSafeGetKieRuntime(kieContainer);
    }

    private static ReleaseId getReleaseId(final KieServices kieServices) {
        return kieServices.newReleaseId("org.kie", "dmn-test-" + UUID.randomUUID(), "1.0");
    }

    private static DMNRuntime typeSafeGetKieRuntime(final KieContainer kieContainer) {
        final DMNRuntime dmnRuntime = kieContainer.newKieSession().getKieRuntime(DMNRuntime.class);
        ((DMNRuntimeImpl) dmnRuntime).setOption(new RuntimeTypeCheckOption(true));
        return dmnRuntime;
    }

    private void start() throws IOException {

        int port = 50051;

        server = ServerBuilder.forPort(port)
                .addService(new GameEngineImpl())
                .build()
                .start();

        logger.info("The game server is running on port: " + port);

        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            System.out.println("Shutting down the game server...");
            KieGrpcServer.this.stop();
            System.out.println("Done.");
        }));
    }

    private void stop() {
        if (server != null) {
            server.shutdown();
        }
    }

    private void blockUntilShutdown() throws InterruptedException {
        if (server != null) {
            server.awaitTermination();
        }
    }

    static class GameEngineImpl extends GameEngineGrpc.GameEngineImplBase {

        @Override
        public void process(final GameInput input,
                            final StreamObserver<GameOutput> responseObserver) {

            final DMNRuntime runtime = createRuntime(input.getModelName() + ".dmn");
            final DMNModel dmnModel = runtime.getModel(input.getModelNamespace(), input.getModelName());

            final DMNResult dmnResult = runtime.evaluateAll(dmnModel, makeDMNContext(input));
            final GameOutput build = makeOutput(dmnResult.getContext().getAll());

            responseObserver.onNext(build);
            responseObserver.onCompleted();
        }
    }

    private static GameOutput makeOutput(final Map<String, Object> map) {
        final String message = getMessage(map);
        final String item = getItem(map);
        final String modelName = getModelName(map);
        final String modelNamespace = getModelNamespace(map);
        return GameOutput
                .newBuilder()
                .setMessage(message)
                .setItem(item)
                .setModelName(modelName)
                .setModelNamespace(modelNamespace)
                .build();
    }

    private static DMNContext makeDMNContext(final GameInput input) {
        final DMNContext context = DMNFactory.newContext();
        context.set("target", input.getTarget());
        context.set("item", input.getItem());
        context.set("action", input.getAction());
        return context;
    }

    private static String getMessage(final Map<String, Object> map) {
        return getStringValue(map, "message");
    }

    private static String getItem(final Map<String, Object> map) {
        return getStringValue(map, "obtained item");
    }

    private static String getModelName(final Map<?, ?> map) {
        return getModel(map).map(m -> getStringValue(m, "modelName")).orElse("");
    }

    private static String getModelNamespace(final Map<?, ?> map) {
        return getModel(map).map(m -> getStringValue(m, "modelNamespace")).orElse("");
    }

    private static Optional<Map<?, ?>> getModel(final Map<?, ?> map) {
        return Optional.ofNullable(map.get("model")).map(e -> (Map<?, ?>) e);
    }

    private static String getStringValue(final Map<?, ?> map,
                                         final String key) {
        return Optional
                .ofNullable(map.get(key))
                .map(v -> (String) v)
                .orElse("");
    }
}
