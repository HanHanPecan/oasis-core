//! Transaction client generating macros.

/// Create a transaction client for a given API.
///
/// # Examples
///
/// This macro should be invoked using a concrete API generated by `runtime_api` as
/// follows:
/// ```rust,ignore
/// with_api! {
///     create_txn_api_client!(MyClient, api);
/// }
/// ```
///
/// In this example, the generated client type will be called `MyClient`. The API
/// definitions will passed as the last argument as defined by the `api` token.
#[macro_export]
macro_rules! create_txn_api_client {
    (
        $name:ident,

        $(
            pub fn $method_name: ident ( $request_type: ty ) -> $response_type: ty ;
        )*
    ) => {
        pub struct $name {
            txn_client: $crate::TxnClient,
        }

        impl $name {
            /// Create new client instance.
            pub fn new(txn_client: $crate::TxnClient) -> Self {
                Self {
                    txn_client,
                }
            }

            /// The underlying transaction client.
            pub fn txn_client(&self) -> &$crate::TxnClient {
                &self.txn_client
            }

            // Generate methods.
            $(
                pub async fn $method_name(
                    &self,
                    arguments: $request_type
                ) -> Result<$response_type, $crate::transaction::macros::Error> {
                    Ok(self.txn_client.call(stringify!($method_name), arguments).await?)
                }
            )*
        }
    };
}

// Re-exported for use in macros.
pub use anyhow::Error;