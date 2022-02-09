package grpcuiredoc

const cssContent = `
/**
 * Theme: nhatthm/grpcui-redoc
 */

@import url('https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;700&family=Roboto+Mono:wght@300;400;700&family=Roboto:wght@300;400;700&display=swap');

h1 {
	font-family: Montserrat, sans-serif !important;
}

img.grpc-logo {
	margin-right: 0;
}

.heading h2 {
	position: relative;
	visibility: collapse;
}

.target {
	visibility: visible;
	position: absolute;
    left: 0;
    font-style: normal;
    font-family: Montserrat, sans-serif;
}

.target::after {
    display: block;
    font-size: 11px;
    color: rgba(250, 250, 250, 0.7);
	margin-top: 4px;
	content: "powered by nhatthm/grpcui-redoc";
}

#grpc-request-response .grpc-tabcontent.ui-tabs-panel h3 {
	border-bottom: 1px solid rgba(38, 50, 56, 0.3);
    margin: 1em 0px;
    color: rgba(38, 50, 56, 0.5);
    font-weight: normal;
    text-transform: uppercase;
    font-size: 0.929em;
    line-height: 20px;
}

#grpc-request-tab {
    display: flex;
    flex-direction: column;
}

#grpc-request-tab > h3:first-child {
	order: 2;
}

#grpc-request-tab #grpc-request-metadata {
	order: 3;
}

#grpc-request-tab > h3:nth-child(5) {
	order: 4;
}

#grpc-request-tab #grpc-request-timeout {
	order: 5;
}

#grpc-request-tab > button {
	order: 6;
	width: 100px;
	margin-top: 10px;
	cursor: pointer;
}

#grpc-request-metadata-form input {
	padding: 6px;
}

#grpc-form .grpc-request-table button.delete,
#grpc-history-list button.delete {
	background: rgb(204, 51, 51);
	border: 0;
}

#grpc-form .grpc-request-table button.add {
	background: rgb(47, 129, 50);
	border: 0;
}

.grpc-desc span.grpc-form-label {
	display: inline-block;
	float: none;
	font-size: 16px;
	line-height: 34px;
}

#grpc-service, #grpc-method {
    font-size: 16px;
    font-family: Montserrat, sans-serif;
    -webkit-font-smoothing: antialiased;
    text-rendering: optimizespeed !important;
    padding: 6px 26px 6px 6px;
    height: 34px;
    -webkit-appearance: none;
    -moz-appearance: none;
    background: transparent;
    background-image: url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>");
    background-repeat: no-repeat;
    background-position-x: 100%;
    background-position-y: 8px;
    border: 1px solid #11181A;
    border-width: 0 0 1px;
    outline: none;
}

#grpc-service option, #grpc-method option {
	font-size: 14px;
}

.ui-widget {
    font-family: Roboto, sans-serif;
    font-weight: 400;
    line-height: 1.5em;
}

#grpc-request-form th {
	font-family: Montserrat, sans-serif;
    font-weight: 400;
    font-size: 22px;
    line-height: 35px;
    color: black;
	padding-bottom: 16px;
}

#grpc-request-form td {
	vertical-align: top;
}

#grpc-request-form div.input_container {
	padding-left: 0;
	border: 0;
}

#grpc-request-form > div.input_container {
	padding-top: 6px;
}

#grpc-request-form div.input_container div.input_container {
	padding-top: 0;
}

#grpc-request-form table {
	border-spacing: 0px;
	border-collapse: separate;
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td {
	padding-top: 0;
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td.toggle_presence,
#grpc-request-form div.input_container tr.message_field:first-of-type > td.toggle_presence {
	padding-top: 8px !important;
}

#grpc-request-form table > tr:nth-child(2) > td:first-child,
#grpc-request-form div.input_container tr.message_field:first-of-type > td:first-child,
#grpc-request-form div.input_container tr:last-child > td:first-child {
	border-left-width: 0px !important;
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
}

#grpc-request-form table > tr:nth-child(2) > td:first-child,
#grpc-request-form div.input_container tr.message_field:first-of-type > td:first-child {
	background-image: linear-gradient(transparent 0%, transparent 8px, rgb(124, 124, 187) 8px, rgb(124, 124, 187) 100%);
}

#grpc-request-form div.input_container tr:last-child > td:first-child {
	background-image: linear-gradient(rgb(124, 124, 187) 0%, rgb(124, 124, 187) 24px, transparent 24px, transparent 100%);
}

#grpc-request-form div.input_container tr.message_field > td:first-child {
	border-left-width: 0px;
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
	border-left: 1px solid rgb(124, 124, 187);
	padding-left: 20px;
	position: relative;
}

#grpc-request-form tr.message_field > td:first-child strong::before {
    content: "";
    display: inline-block;
    vertical-align: middle;
    width: 10px;
    height: 1px;
    background: rgb(124, 124, 187);
	position: absolute;
	top: 24px;
	left: 0;
}

#grpc-request-form tr.message_field > td:first-child strong::after {
    content: "";
    display: inline-block;
    vertical-align: middle;
    width: 1px;
    background: rgb(124, 124, 187);
    height: 7px;
	position: absolute;
	top: 21px;
	left: 10px;
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td:first-child strong::before,
#grpc-request-form div.input_container tr.message_field:first-of-type > td:first-child strong::before {
	top: 8px;
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td:first-child strong::after,
#grpc-request-form div.input_container tr.message_field:first-of-type > td:first-child strong::after {
	top: 5px;
}

#grpc-request-form td.name {
	text-align: left;
	font-family: Roboto, sans-serif;
    color: rgb(102, 102, 102);
}

#grpc-request-form td.name strong {
	font-family: "Roboto Mono", "Courier New", Courier, monospace;
    font-weight: 400;
    font-size: 13px;
	color: rgb(51, 51, 51);
}

#grpc-request-form td.toggle_presence {
	padding-top: 24px !important;
}

#grpc-history-list .history-item-header .history-item-method {
	font-family: Montserrat, sans-serif;
    font-weight: 400;
}

#grpc-request-form select {
	-webkit-appearance: none;
	-moz-appearance: none;
	background: transparent;
	background-image: url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>");
	background-repeat: no-repeat;
	background-position-x: 100%;
	background-position-y: 7px;
	padding: 6px 16px 6px 6px;
}

#grpc-request-form select:disabled {
	color: rgb(204,204,204);
	background-image: url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon fill='rgb(204,204,204)' stroke='rgb(204,204,204)' points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>");
}

#grpc-request-form select:focus,
#grpc-request-form select:focus-visible,
#grpc-request-form select:active,
#grpc-request-form select {
	border-width: 0 0 1px;
	outline: none;
}

#grpc-request-form textarea {
	font-family: "Roboto Mono", "Courier New", Courier, monospace;
	min-height: 33px;
	min-width: 400px;
	max-width: 400px;
	padding: 6px;
}

#grpc-request-form input[type=text] {
	width: 400px;
	padding: 6px;
}

#grpc-request-form .null {
	line-height: 1.5em;
	padding: 6px 6px 6px 0;
}

#grpc-response-error {
	color: rgb(212, 31, 28);
	background: rgba(212, 31, 28, 0.07);
	font-family: Roboto, sans-serif;
    font-weight: 400;
    line-height: 1.5em;
	padding: 10px;
	font-size: 100%;
}

#grpc-response-error-summary {
	font-weight: 700;
}

#grpc-response-error #grpc-response-error-num {
	color: inherit;
}

#grpc-response-error #grpc-response-error-msg {
	font-family: Roboto, sans-serif;
	font-weight: 300;
}

#grpc-history-list .history-item-header .history-item-result.error {
    color: rgb(212, 31, 28);
}
`
