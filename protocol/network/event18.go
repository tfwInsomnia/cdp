// +build !go1.9

// Code generated by cdpgen. DO NOT EDIT.

package network

import (
	"github.com/mafredri/cdp/protocol"
	"github.com/mafredri/cdp/rpcc"
)

// RequestWillBeSentReply is the reply for RequestWillBeSent events.
type RequestWillBeSentReply struct {
	RequestID        RequestID                  `json:"requestId"`                  // Request identifier.
	LoaderID         LoaderID                   `json:"loaderId"`                   // Loader identifier. Empty string if the request is fetched form worker.
	DocumentURL      string                     `json:"documentURL"`                // URL of the document this request is loaded for.
	Request          Request                    `json:"request"`                    // Request data.
	Timestamp        MonotonicTime              `json:"timestamp"`                  // Timestamp.
	WallTime         TimeSinceEpoch             `json:"wallTime"`                   // Timestamp.
	Initiator        Initiator                  `json:"initiator"`                  // Request initiator.
	RedirectResponse *Response                  `json:"redirectResponse,omitempty"` // Redirect response data.
	Type             *protocol.PageResourceType `json:"type,omitempty"`             // Type of this resource.
	FrameID          *protocol.PageFrameID      `json:"frameId,omitempty"`          // Frame identifier.
}

// RequestServedFromCacheClient is a client for RequestServedFromCache events. Fired if request ended up loading from cache.
type RequestServedFromCacheClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*RequestServedFromCacheReply, error)
	rpcc.Stream
}

// ResponseReceivedReply is the reply for ResponseReceived events.
type ResponseReceivedReply struct {
	RequestID RequestID                 `json:"requestId"`         // Request identifier.
	LoaderID  LoaderID                  `json:"loaderId"`          // Loader identifier. Empty string if the request is fetched form worker.
	Timestamp MonotonicTime             `json:"timestamp"`         // Timestamp.
	Type      protocol.PageResourceType `json:"type"`              // Resource type.
	Response  Response                  `json:"response"`          // Response data.
	FrameID   *protocol.PageFrameID     `json:"frameId,omitempty"` // Frame identifier.
}

// DataReceivedClient is a client for DataReceived events. Fired when data chunk was received over the network.
type DataReceivedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DataReceivedReply, error)
	rpcc.Stream
}

// LoadingFailedReply is the reply for LoadingFailed events.
type LoadingFailedReply struct {
	RequestID     RequestID                 `json:"requestId"`               // Request identifier.
	Timestamp     MonotonicTime             `json:"timestamp"`               // Timestamp.
	Type          protocol.PageResourceType `json:"type"`                    // Resource type.
	ErrorText     string                    `json:"errorText"`               // User friendly error message.
	Canceled      *bool                     `json:"canceled,omitempty"`      // True if loading was canceled.
	BlockedReason BlockedReason             `json:"blockedReason,omitempty"` // The reason why loading was blocked, if any.
}

// WebSocketWillSendHandshakeRequestClient is a client for WebSocketWillSendHandshakeRequest events. Fired when WebSocket is about to initiate handshake.
type WebSocketWillSendHandshakeRequestClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WebSocketWillSendHandshakeRequestReply, error)
	rpcc.Stream
}

// RequestInterceptedReply is the reply for RequestIntercepted events.
type RequestInterceptedReply struct {
	InterceptionID     InterceptionID            `json:"interceptionId"`               // Each request the page makes will have a unique id, however if any redirects are encountered while processing that fetch, they will be reported with the same id as the original fetch. Likewise if HTTP authentication is needed then the same fetch id will be used.
	Request            Request                   `json:"request"`                      //
	ResourceType       protocol.PageResourceType `json:"resourceType"`                 // How the requested resource will be used.
	RedirectHeaders    Headers                   `json:"redirectHeaders,omitempty"`    // HTTP response headers, only sent if a redirect was intercepted.
	RedirectStatusCode *int                      `json:"redirectStatusCode,omitempty"` // HTTP response code, only sent if a redirect was intercepted.
	RedirectURL        *string                   `json:"redirectUrl,omitempty"`        // Redirect location, only sent if a redirect was intercepted.
	AuthChallenge      *AuthChallenge            `json:"authChallenge,omitempty"`      // Details of the Authorization Challenge encountered. If this is set then continueInterceptedRequest must contain an authChallengeResponse.
}