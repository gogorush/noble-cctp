/*
 * Copyright (c) 2023, © Circle Internet Financial, LTD.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cli

import (
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/gogorush/noble-cctp/x/cctp/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

func CmdDepositForBurn() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit-for-burn [amount] [destination-domain] [mint-recipient] [burn-token]",
		Short: "Deposit For Burn",
		Long:  "Broadcast a transaction that deposits for burn to a provided domain.",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			amount, ok := math.NewIntFromString(args[0])
			if !ok {
				return sdkerrors.Wrapf(types.ErrInvalidAmount, "invalid amount")
			}

			destinationDomain, err := strconv.ParseUint(args[1], types.BaseTen, types.DomainBitLen)
			if err != nil {
				return err
			}

			mintRecipient, err := parseAddress(args[2])
			if err != nil {
				return fmt.Errorf("invalid mint recipient: %w", err)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositForBurn(
				clientCtx.GetFromAddress().String(),
				amount,
				uint32(destinationDomain),
				mintRecipient,
				args[3],
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
