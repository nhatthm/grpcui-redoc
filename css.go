// Code generated by resources/scripts/generate.sh. DO NOT EDIT.
// source: resources/web/style.css

package grpcuiredoc

const cssContent = `
/**
 * Theme: nhatthm/grpcui-redoc
 */

/* stylelint-disable selector-class-pattern */
/* stylelint-disable no-descending-specificity */

@import "https://fonts.googleapis.com/css2?family=Montserrat:wght@300;400;700&family=Roboto+Mono:wght@300;400;700&family=Roboto:wght@300;400;700&display=swap";

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
    color: rgb(250 250 250 / 70%);
    margin-top: 4px;
    content: "powered by nhatthm/grpcui-redoc";
}

#grpc-request-tab > h3:first-child {
    order: 2;
}

#grpc-request-tab {
    display: flex;
    flex-direction: column;
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

#grpc-history-list button.delete,
#grpc-form .grpc-request-table button.delete {
    background: #c33;
    border: 0;
}

#grpc-form .grpc-request-table button.add {
    background: #2f8132;
    border: 0;
}

.grpc-desc span.grpc-form-label {
    display: inline-block;
    float: none;
    font-size: 16px;
    line-height: 34px;
}

#grpc-service,
#grpc-method {
    font-size: 16px;
    font-family: Montserrat, sans-serif;
    -webkit-font-smoothing: antialiased;
    text-rendering: optimizespeed !important;
    padding: 6px 26px 6px 6px;
    height: 34px;
    -webkit-appearance: none;
       -moz-appearance: none;
            appearance: none;
    background: transparent url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>") no-repeat;
    background-position-x: 100%;
    background-position-y: 8px;
    border-style: solid;
    border-color: #11181a;
    border-width: 0 0 1px;
    outline: none;
}

#grpc-service option,
#grpc-method option {
    font-size: 14px;
}

.ui-widget {
    font-family: Roboto, sans-serif;
    font-weight: 400;
    line-height: 1.5em;
}

#grpc-request-response .grpc-tabcontent.ui-tabs-panel h3 {
    border-bottom: 1px solid rgb(38 50 56 / 30%);
    margin: 1em 0;
    color: rgb(38 50 56 / 50%);
    font-weight: normal;
    text-transform: uppercase;
    font-size: 0.929em;
    line-height: 20px;
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
    border-spacing: 0;
    border-collapse: separate;
}

#grpc-request-form td.toggle_presence {
    padding-top: 24px !important;
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td.toggle_presence,
#grpc-request-form div.input_container tr.message_field:first-of-type > td.toggle_presence {
    padding-top: 8px !important;
}

#grpc-request-form table > tr:nth-child(2) > td:first-child {
    border-left-width: 0 !important;
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
    background-image: linear-gradient(transparent 0%, transparent 8px, rgb(124 124 187) 8px, rgb(124 124 187) 100%);
}

#grpc-request-form > div.input_container > table > tr:nth-child(2) > td {
    padding-top: 0;
}

#grpc-request-form div.input_container tr:last-child > td:first-child {
    border-left-width: 0 !important;
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
    background-image: linear-gradient(rgb(124 124 187) 0%, rgb(124 124 187) 24px, transparent 24px, transparent 100%);
}

#grpc-request-form div.input_container tr.message_field > td:first-child {
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
    border-left: 1px solid #7c7cbb;
    padding-left: 20px;
    position: relative;
}

#grpc-request-form div.input_container tr.message_field:first-of-type > td:first-child {
    border-left-width: 0 !important;
    background-position: left top;
    background-repeat: no-repeat;
    background-size: 1px 100%;
    background-image: linear-gradient(transparent 0%, transparent 8px, rgb(124 124 187) 8px, rgb(124 124 187) 100%);
}

#grpc-request-form tr.message_field > td:first-child strong::before {
    content: "";
    display: inline-block;
    vertical-align: middle;
    width: 10px;
    height: 1px;
    background: #7c7cbb;
    position: absolute;
    top: 24px;
    left: 0;
}

#grpc-request-form tr.message_field > td:first-child strong::after {
    content: "";
    display: inline-block;
    vertical-align: middle;
    width: 1px;
    background: #7c7cbb;
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
    color: #666;
}

#grpc-request-form td.name strong {
    font-family: "Roboto Mono", "Courier New", Courier, monospace;
    font-weight: 400;
    font-size: 13px;
    color: #333;
}

#grpc-history-list .history-item-header .history-item-method {
    font-family: Montserrat, sans-serif;
    font-weight: 400;
}

#grpc-request-form select {
    -webkit-appearance: none;
       -moz-appearance: none;
            appearance: none;
    background: transparent url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>") no-repeat;
    background-position-x: 100%;
    background-position-y: 7px;
    padding: 6px 16px 6px 6px;
}

#grpc-request-form select,
#grpc-request-form select:focus,
#grpc-request-form select:focus-visible,
#grpc-request-form select:active {
    border-width: 0 0 1px;
    outline: none;
}

#grpc-request-form select:disabled {
    color: #ccc;
    background-image: url("data:image/svg+xml;utf8,<svg version='1.1' viewBox='0 0 32 48' x='0' xmlns='http://www.w3.org/2000/svg' y='0' aria-hidden='true' style='margin-right: -25px;'><polygon fill='rgb(204,204,204)' stroke='rgb(204,204,204)' points='17.3 8.3 12 13.6 6.7 8.3 5.3 9.7 12 16.4 18.7 9.7 '></polygon></svg>");
}

#grpc-request-form textarea {
    font-family: "Roboto Mono", "Courier New", Courier, monospace;
    min-height: 33px;
    min-width: 400px;
    max-width: 400px;
    padding: 6px;
}

#grpc-request-form input[type="text"] {
    width: 400px;
    padding: 6px;
}

#grpc-request-form .null {
    line-height: 1.5em;
    padding: 6px 6px 6px 0;
}

#grpc-response-error {
    color: #d41f1c;
    background: rgb(212 31 28 / 7%);
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
    color: #d41f1c;
}

.grpc-desc,
.grpc-desc > tbody,
.grpc-desc > tbody > tr > td {
    display: block;
}

.grpc-desc > tbody > tr {
    display: inline-block;
    position: relative !important;
    overflow: visible;
}

.grpc-desc pre {
    font-family: "Roboto Mono", "Courier New", Courier, monospace;
    font-size: 12px;
}

#grpc-descriptions {
    position: absolute;
    top: 0;
    left: 100%;
    overflow: visible;
    z-index: 1000;
    margin-left: 36px;
}

#grpc-descriptions-toggle {
    left: -36px;
    cursor: pointer;
}

#grpc-service-description {
    margin: 0;
    background-color: rgba(255 255 255 / 85%);
    background-clip: padding-box;
    border-width: 1px 1px 0;
    border-style: solid;
    border-color: rgba(0 0 0 / 10%);
    border-radius: 0.25rem 0.25rem 0 0;
    box-shadow: rgb(0 0 0 / 10%) 0 -0.25rem 0.75rem;
    position: relative;
    z-index: 2;
}

#grpc-service-description::after {
    content: " ";
    height: 0.5em;
    top: 100%;
    left: 0;
    background: #fff;
    position: absolute;
    width: 100%;
    z-index: 4;
}

#grpc-method-description {
    margin-left: 0;
    background-color: rgba(255 255 255 / 85%);
    background-clip: padding-box;
    border-width: 0 1px;
    border-style: solid;
    border-color: rgba(0 0 0 / 10%);
    border-radius: 0;
    box-shadow: rgb(0 0 0 / 10%) 0 0 0.75rem;
    position: relative;
    z-index: 1;
}

#grpc-method-description::before {
    content: " ";
    position: absolute;
    top: -0.5em;
    left: 0;
    z-index: 2;
    width: 100%;
    height: 0.5em;
    background: rgba(255 255 255 / 85%);
}

#grpc-method-description::after {
    content: " ";
    position: absolute;
    z-index: 2;
    top: 100%;
    left: 0;
    width: 100%;
    height: 0.5em;
    background: rgba(255 255 255 / 85%);
}

#grpc-service-description-end {
    margin-left: 0;
    background-color: rgba(255 255 255 / 85%);
    background-clip: padding-box;
    border-width: 0 1px 1px;
    border-style: solid;
    border-color: rgba(0 0 0 / 10%);
    border-radius: 0 0 0.25rem 0.25rem;
    box-shadow: rgb(0 0 0 / 10%) 0 0.25rem 0.75rem;
    position: relative;
    z-index: 3;
}

#grpc-service-description-end::before {
    content: " ";
    position: absolute;
    z-index: 4;
    top: -0.5em;
    left: 0;
    width: 100%;
    height: 0.5em;
    background: rgba(255 255 255 / 85%);
}

#grpc-service-description,
#grpc-service-description-end {
    color: #333;
}
`
