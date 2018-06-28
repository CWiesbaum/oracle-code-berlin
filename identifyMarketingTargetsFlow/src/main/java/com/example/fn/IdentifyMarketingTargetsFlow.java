package com.example.fn;

import com.fnproject.fn.api.flow.Flow;
import com.fnproject.fn.api.flow.FlowFuture;
import com.fnproject.fn.api.flow.Flows;
import com.fnproject.fn.api.flow.HttpResponse;

import static com.fnproject.fn.api.flow.HttpMethod.POST;
import static com.fnproject.fn.api.Headers.emptyHeaders;

public class IdentifyMarketingTargetsFlow {
	private static final String SEARCH_PRODUCT_FUNCTION = "./searchProduct";
	private static final String EXTRACT_EMAIL_FUNCTION = "./extractEmail";

    public String identifyMarketingTargets(String input) {
    	Flow fl = Flows.currentFlow();
    	
    	FlowFuture<byte[]> emailList = fl.invokeFunction(SEARCH_PRODUCT_FUNCTION, POST, emptyHeaders(), input.getBytes())
    			.thenApply(HttpResponse::getBodyAsBytes)
    			.thenCompose(identifiedTargets -> fl.invokeFunction(EXTRACT_EMAIL_FUNCTION, POST, emptyHeaders(), identifiedTargets))
    			.thenApply(HttpResponse::getBodyAsBytes);
    			;
    	
        return new String(emailList.get());
    }

}