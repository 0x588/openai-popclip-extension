#!/bin/bash
if [[ $(uname -m) == 'arm64' ]]
then
./openai-cmd-arm64 -k="$POPCLIP_OPTION_APIKEY" -u="$POPCLIP_OPTION_PROXY"  -p="$POPCLIP_OPTION_PROMPT3" -c="$POPCLIP_TEXT" -m="$POPCLIP_OPTION_MODEL" -s=$POPCLIP_OPTION_SHOW
else
./openai-cmd -k="$POPCLIP_OPTION_APIKEY" -u="$POPCLIP_OPTION_PROXY"  -p="$POPCLIP_OPTION_PROMPT3" -c="$POPCLIP_TEXT" -m="$POPCLIP_OPTION_MODEL" -s=$POPCLIP_OPTION_SHOW
fi